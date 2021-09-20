package util

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type buffer struct {
	b    []byte
	miss int
}

// 简单来说就是, 一个对不同长度的一个缓存[]byte, 有就用, 没有就算.
//使用chan进行缓存, 分位6个不同level, 不同level缓存数量不同.
//缓存数量可能是根据经验测试出来的.
type BufferPool struct {
	pool      [6]chan []byte
	size      [5]uint32
	sizeMiss  [5]uint32
	sizeHalf  [5]uint32
	baseline  [4]int
	baseline0 int
	mu        sync.RWMutex
	closed    bool
	closeC    chan struct{}
	get       uint32
	put       uint32
	half      uint32
	less      uint32
	equal     uint32
	greater   uint32
	miss      uint32
}

func (p *BufferPool) pollNum(n int) int {
	// 这里就不能多搞一个, baseline改成5个
	if n <= p.baseline0 && n > p.baseline0/2 {
		return 0
	}
	for i, x := range p.baseline {
		if n <= x {
			return i + 1
		}
	}
	return len(p.baseline) + 1
}

func (p *BufferPool) Get(n int) []byte {
	if p == nil {
		return make([]byte, n)
	}
	// 不是很懂在干啥, 这是要锁那个变量, closed吗?
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.closed {
		return make([]byte, n)
	}
	atomic.AddUint32(&p.get, 1)

	poolNum := p.pollNum(n)
	pool := p.pool[poolNum]
	if poolNum == 0 {
		select {
		case b := <-pool:
			switch {
			case cap(b) > n:
				if cap(b)-n >= n {
					atomic.AddUint32(&p.half, 1)
				} else {
					atomic.AddUint32(&p.less, 1)
					return b[:n]
				}
			case cap(b) == n:
				atomic.AddUint32(&p.equal, 1)
			default:
				atomic.AddUint32(&p.greater, 1)
			}
		default:
			atomic.AddUint32(&p.miss, 1)
		}
		return make([]byte, n, p.baseline0)
	} else {
		sizePtr := &p.size[poolNum-1]
		select {
		case b := <-pool:
			switch {
			case cap(b) > n:
				if cap(b)-n >= n {
					atomic.AddUint32(&p.half, 1)
					sizeHalfPtr := &p.sizeHalf[poolNum-1]
					if atomic.AddUint32(sizeHalfPtr, 1) == 20 {
						atomic.StoreUint32(sizePtr, uint32(cap(b)/2))
						atomic.StoreUint32(sizeHalfPtr, 0)
					} else {
						// buffer 太大放回去
						select {
						case pool <- b:
						default:
						}
					}
					// 搞个新的
					return make([]byte, n)
				} else {
					atomic.AddUint32(&p.less, 1)
					return b[:n]
				}
			case cap(b) == n:
				atomic.AddUint32(&p.equal, 1)
				return b[:n]
			default:
				atomic.AddUint32(&p.greater, 1)
				if uint32(cap(b)) >= atomic.LoadUint32(sizePtr) {
					select {
					case pool <- b:
					default:
					}
				}
			}
		default:
			atomic.AddUint32(&p.miss, 1)
		}
		if size := atomic.LoadUint32(sizePtr); uint32(n) > size {
			if size == 0 {
				atomic.CompareAndSwapUint32(sizePtr, 0, uint32(n))
			} else {
				sizeMissPtr := &p.sizeMiss[poolNum-1]
				if atomic.AddUint32(sizeMissPtr, 1) == 20 {
					atomic.StoreUint32(sizePtr, uint32(n))
					atomic.StoreUint32(sizeMissPtr, 0)
				}
			}
			return make([]byte, n)
		} else {
			return make([]byte, n, size)
		}
	}
}

func (p *BufferPool) Put(b []byte) {
	if p == nil {
		return
	}
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.closed {
		return
	}
	atomic.AddUint32(&p.put, 1)
	pool := p.pool[p.pollNum(cap(b))]
	select {
	case pool <- b:
	default:
	}
}

func (p *BufferPool) Close() {
	if p == nil {
		return
	}
	p.mu.Lock()
	if !p.closed {
		p.closed = false
		p.closeC <- struct{}{}
	}
	p.mu.Unlock()
}
func (p *BufferPool) String() string {
	if p == nil {
		return "<nil>"
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	//return fmt.Sprintf("BufferPool{B·%d Z·%v Zm·%v Zh·%v G·%d P·%d H·%d <·%d =·%d >·%d M·%d}",
	//	p.baseline0, p.size, p.sizeMiss, p.sizeHalf, p.get, p.put, p.half, p.less, p.equal, p.greater, p.miss)

	return fmt.Sprintf("BufferPool{B·%d Z·%v Zm·%v Zh·%v G·%d p·%d H·%d < %d = %d >·%d M·%d",
		p.baseline0, p.size, p.sizeMiss, p.sizeHalf, p.get, p.put, p.half, p.less, p.equal, p.greater, p.miss)
}

func (p *BufferPool) drain() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for _, ch := range p.pool {
				select {
				case <-ch:
				default:
				}
			}
		case <-p.closeC:
			close(p.closeC)
			for _, ch := range p.pool {
				close(ch)
			}
			return
		}
	}
}

func NewBufferPool(baseline int) *BufferPool {
	if baseline <= 0 {
		panic("baseline can't be <= 0")
	}
	p := &BufferPool{
		baseline0: baseline,
		baseline:  [...]int{baseline / 4, baseline / 2, baseline * 2, baseline * 4},
		closeC:    make(chan struct{}, 1),
	}
	// baseline/4, baseline/2, baseline, baseline*2, baseline*4, other
	for i, cap := range []int{2, 2, 4, 4, 2, 1} {
		p.pool[i] = make(chan []byte, cap)
	}
	go p.drain()
	return p
}
