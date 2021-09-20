package cache

import (
	"github.com/zoroqi/transcribe-code/goleveldb/leveldb/util"
	"sync"
	"sync/atomic"
	"unsafe"
)

// 具体的cache
type Cache struct {
	mu     sync.RWMutex
	mHead  unsafe.Pointer
	nodes  int32
	size   int32 // 数量
	cacher Cacher
	closed bool
}

/**

Cache



*/

func NewCache(cacher Cacher) *Cache {
	h := &mNode{
		buckets:         make([]unsafe.Pointer, mInitalSize),
		mask:            mInitalSize - 1,
		growThreshold:   int32(mInitalSize * mOverflowThreshold),
		shrinkThreshold: 0,
	}
	for i := range h.buckets {
		h.buckets[i] = unsafe.Pointer(&mBucket{})
	}
	r := &Cache{
		mHead:  unsafe.Pointer(h),
		cacher: cacher,
	}
	return r
}
func (r *Cache) getBucket(hash uint32) (*mNode, *mBucket) {
	h := (*mNode)(atomic.LoadPointer(&r.mHead))
	i := hash & h.mask
	b := (*mBucket)(atomic.LoadPointer(&h.buckets[i]))
	if b == nil {
		b = h.initBucket(i)
	}
	return h, b
}

func (r *Cache) delete(n *Node) bool {
	for {
		h, b := r.getBucket(n.hash)
		done, deleted := b.delete(r, h, n.hash, n.ns, n.key)
		if done {
			return deleted
		}
	}
}

func (r *Cache) Nodes() int {
	return int(atomic.LoadInt32(&r.nodes))
}

func (r *Cache) Size() int {
	return int(atomic.LoadInt32(&r.size))
}

func (r *Cache) Capacity() int {
	if r.cacher == nil {
		return 0
	}
	return r.cacher.Capacity()
}

func (r *Cache) SetCapacity(capacity int) {
	if r.cacher != nil {
		r.cacher.SetCapacity(capacity)
	}
}

// 单独起一个名字好些
type setFunc func() (size int, value Value)

func (r *Cache) Get(ns, key uint64, setFunc setFunc) *Handle {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.closed {
		return nil
	}

	hash := murmur32(ns, key, 0xf00)
	for {
		h, b := r.getBucket(hash)
		// 内部创建好(申请)空间, 返回空间地址, 然后再赋值.
		done, _, n := b.get(r, h, hash, ns, key, setFunc == nil)
		if done {
			if n != nil {
				n.mu.Lock()
				if n.value == nil {
					if setFunc == nil {
						n.mu.Unlock()
						n.unref() // 可能会删除节点
						return nil
					}
					// 不理解这个size是啥意思
					n.size, n.value = setFunc()
					if n.value == nil {
						n.size = 0
						n.mu.Unlock()
						n.unref()
						return nil
					}
					atomic.AddInt32(&r.size, int32(n.size))
				}
				n.mu.Unlock()
				if r.cacher != nil {
					r.cacher.Promote(n)
				}
				return &Handle{unsafe.Pointer(n)}
			}
			break
		}
	}
	return nil
}

func (r *Cache) Delete(ns, key uint64, onDel func()) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.closed {
		return false
	}
	hash := murmur32(ns, key, 0xf00)
	for {
		h, b := r.getBucket(hash)
		// 找到接口
		done, _, n := b.get(r, h, h, ns, key, true)
		if done {
			if n != nil {
				if onDel != nil {
					n.mu.Lock()
					n.onDel = append(n.onDel, onDel)
					n.mu.Unlock()
				}
				if r.cacher != nil {
					r.cacher.Ban(n)
				}
				n.unref()
			}
		}
	}
	if onDel != nil {
		onDel()
	}
	return false
}

type Value interface{}

type Node struct {
	r         *Cache
	hash      uint32
	ns, key   uint64
	mu        sync.Mutex
	size      int
	value     Value
	ref       int32
	onDel     []func() // 删除的监听事件
	CacheData unsafe.Pointer
}

func (n *Node) Key() uint64 {
	return n.key
}

func (n *Node) Ns() uint64 {
	return n.ns
}

func (n *Node) Size() int {
	return n.size
}

func (n *Node) Value() Value {
	return n.value
}

func (n *Node) Ref() int32 {
	return atomic.LoadInt32(&n.ref)
}

func (n *Node) GetHandle() *Handle {
	if atomic.AddInt32(&n.ref, 1) <= 1 {
		panic("BUG: Node.GetHandle on zero ref")
	}
	return &Handle{unsafe.Pointer(n)}
}

// 引用计数, 引用为零执行删除操作
func (n *Node) unref() {
	if atomic.AddInt32(&n.ref, -1) == 0 {
		// 真正的删除
		n.r.delete(n)
	}
}

// 会验证closed状态
func (n *Node) unrefLocked() {
	if atomic.AddInt32(&n.ref, -1) == 0 {
		n.r.mu.RLock()
		if !n.r.closed {
			n.r.delete(n)
		}
		n.r.mu.RUnlock()
	}
}

type Handle struct {
	n unsafe.Pointer
}

func (h *Handle) Value() Value {
	n := (*Node)(atomic.LoadPointer(&h.n))
	if n != nil {
		return n.value
	}
	return nil
}

func (h *Handle) Release() {
	nPtr := atomic.LoadPointer(&h.n)
	if nPtr != nil && atomic.CompareAndSwapPointer(&h.n, nPtr, nil) {
		// 地址强转成node
		n := (*Node)(nPtr)
		n.unrefLocked()
	}
}

// 定义cache操作如何操作等处理
type Cacher interface {
	Capacity() int
	SetCapacity(capacity int)

	Promote(n *Node)
	// 驱逐
	Ban(n *Node)
	Evict(n *Node)
	// ns 是 namespace
	EvictNs(ns uint64)
	EvictAll()
	Close() error
}

const (
	mInitalSize             = 1 << 4
	mOverflowThreshold      = 1 << 5
	mOverflowGrowThreashold = 1 << 7
)

type mNode struct {
	buckets         []unsafe.Pointer
	mask            uint32
	pred            unsafe.Pointer //*mNode
	resizeInProgess int32
	overflow        int32
	growThreshold   int32
	shrinkThreshold int32
}

type mBucket struct {
	mu     sync.Mutex
	node   []*Node
	frozen bool
}

func (b *mBucket) freeze() []*Node {
	b.mu.Lock()
	defer b.mu.Unlock()
	if !b.frozen {
		// 原子变量竟然没有bool类型
		b.frozen = true
	}
	return b.node
}

func (b *mBucket) get(r *Cache, h *mNode, hash uint32, ns, key uint64, noset bool) (done, added bool, n *Node) {
	b.mu.Lock()
	// 为啥不defer?
	if b.frozen {
		b.mu.Unlock()
		return
	}

	// scan 是否存在
	for _, n := range b.node {
		// ns啥? 超时
		if n.hash == hash && n.ns == ns && n.key == key {
			atomic.AddInt32(&n.ref, 1)
			b.mu.Unlock()
			return true, false, n
		}
	}
	// 不用写入
	if noset {
		b.mu.Unlock()
		return true, false, nil
	}
	// 创建新的node
	n = &Node{
		r:    r,
		hash: hash,
		ns:   ns,
		key:  key,
		ref:  1, // 引用
	}
	// 锁到这里就结束了
	b.node = append(b.node, n)
	bLen := len(b.node)
	b.mu.Unlock()

	grow := atomic.AddInt32(&r.nodes, 1) >= h.growThreshold
	if bLen > mOverflowThreshold {
		grow = grow || atomic.AddInt32(&h.overflow, 1) >= mOverflowGrowThreashold
	}
	// 不懂
	if grow && atomic.CompareAndSwapInt32(&h.resizeInProgess, 0, 1) {
		nhLen := len(h.buckets) << 1
		nh := &mNode{
			buckets:         make([]unsafe.Pointer, nhLen),
			mask:            uint32(nhLen) - 1,
			pred:            unsafe.Pointer(h),
			growThreshold:   int32(nhLen * mOverflowThreshold),
			shrinkThreshold: int32(nhLen >> 1),
		}
		ok := atomic.CompareAndSwapPointer(&r.mHead, unsafe.Pointer(h), unsafe.Pointer(nh))
		if !ok {
			panic("BUG: failed swapping head")
		}
		go nh.initBuckets()
	}
	return true, true, n
}

func (b *mBucket) delete(r *Cache, h *mNode, hash uint32, ns, key uint64) (done, deleted bool) {
	b.mu.Lock()
	if b.frozen {
		b.mu.Unlock()
		return
	}
	// 不理解为啥要这么定义
	var (
		n    *Node
		bLen int
	)

	for i := range b.node {
		n = b.node[i]
		if n.ns == ns && n.key == key {
			if atomic.LoadInt32(&n.ref) == 0 {
				deleted = true
				if n.value != nil {
					if r, ok := n.value.(util.Releaser); ok {
						r.Release()
					}
					n.value = nil
				}
				// 不理解golang 移除为啥这么费劲
				b.node = append(b.node[:i], b.node[i+1:]...)
				bLen = len(b.node)
			}
			break
		}
	}
	b.mu.Unlock()
	if deleted {
		for _, f := range n.onDel {
			f()
		}
		atomic.AddInt32(&r.size, int32(n.size)*-1)
		shrink := atomic.AddInt32(&r.nodes, -1) < h.shrinkThreshold
		if bLen >= mOverflowThreshold {
			atomic.AddInt32(&h.overflow, -1)
		}
		if shrink && len(h.buckets) > mInitalSize && atomic.CompareAndSwapInt32(&h.resizeInProgess, 0, 1) {
			nhLen := len(h.buckets) >> 1
			nh := &mNode{
				buckets:         make([]unsafe.Pointer, nhLen),
				mask:            uint32(nhLen) - 1,
				pred:            unsafe.Pointer(h),
				growThreshold:   int32(nhLen * mOverflowThreshold),
				shrinkThreshold: int32(nhLen >> 1),
			}
			ok := atomic.CompareAndSwapPointer(&r.mHead, unsafe.Pointer(h), unsafe.Pointer(nh))
			if !ok {
				panic("BUG: failed swapping head")
			}
			go nh.initBuckets()
		}
		return true, deleted

	}

}

func (n *mNode) initBuckets() {
	for i := range n.buckets {
		n.initBucket(uint32(i))
	}
	atomic.StorePointer(&n.pred, nil)
}

func (n *mNode) initBucket(i uint32) *mBucket {
	if b := (*mBucket)(atomic.LoadPointer(&n.buckets[i])); b != nil {
		return b
	}

	p := (*mNode)(atomic.LoadPointer(&n.pred))
	if p != nil {
		var node []*Node
		if n.mask > p.mask {
			pb := (*mBucket)(atomic.LoadPointer(&p.buckets[i&p.mask]))
			if pb == nil {
				pb = p.initBucket(i & p.mask)
			}
			m := pb.freeze()
			for _, x := range m {
				if x.hash&n.mask == i {
					node = append(node, x)
				}
			}
		} else {
			pb0 := (*mBucket)(atomic.LoadPointer(&p.buckets[i]))
			if pb0 == nil {
				pb0 = p.initBucket(i)
			}
			pb1 := (*mBucket)(atomic.LoadPointer(&p.buckets[i+uint32(len(n.buckets))]))
			if pb1 == nil {
				pb1 = p.initBucket(i + uint32(len(n.buckets)))
			}
			m0 := pb0.freeze()
			m1 := pb1.freeze()
			node = make([]*Node, 0, len(m0)+len(m1))
			node = append(node, m0...)
			node = append(node, m1...)
		}
		b := &mBucket{node: node}
		if atomic.CompareAndSwapPointer(&n.buckets[i], nil, unsafe.Pointer(b)) {
			if len(node) > mOverflowThreshold {
				atomic.AddInt32(&n.overflow, int32(len(node)-mOverflowThreshold))
			}
			return b
		}
	}
	return (*mBucket)(atomic.LoadPointer(&n.buckets[i]))
}

func murmur32(ns, key uint64, seed uint32) uint32 {
	const (
		m = uint32(0x5bd1e995)
		r = 24
	)
	k1 := uint32(ns >> 32)
	k2 := uint32(ns)
	k3 := uint32(key >> 32)
	k4 := uint32(key)
	k1 *= m
	k1 ^= k1 >> r
	k1 *= m

	k2 *= m
	k2 ^= k2 >> r
	k2 *= m
	k3 *= m
	k3 ^= k3 >> r
	k3 *= m

	k4 *= m
	k4 ^= k4 >> r
	k4 *= m

	h := seed

	h *= m
	h ^= k1
	h *= m
	h ^= k2
	h *= m
	h ^= k3
	h *= m
	h ^= k4
	h ^= h >> 13
	h *= m
	h ^= h >> 15
	return h
}
