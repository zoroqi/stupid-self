package other

import (
	"cmp"
	"fmt"
	"sort"
	"strings"
)

type BPTree[K cmp.Ordered, V any] interface {
	Insert(k K, v V) bool
	Search(k K) (V, bool)
	Delete(k K) bool
	Iterator(func(K, V) error) error
}

func NewBPTree[K cmp.Ordered, T any](m int) BPTree[K, T] {
	if !(4 < m && m < 100) {
		m = 5
	}
	b := bptree[K, T]{m: m}
	return &b
}

type tuple[F, S any] struct {
	f F // first
	s S // second
}

// 如果 leaf 是 true, 则 children == nil, len(keys) == len(values);
// 如果 leaf 是 false, 则 values == nil, len(keys) + 1 == len(children).
// keys 和 children 关系是, keys[i] <= children[i] < keys[i+1].
type bpnode[K, V any] struct {
	leaf     bool
	keys     []K
	values   []V
	children []*bpnode[K, V]
	parent   *bpnode[K, V]
}

type bptree[K cmp.Ordered, V any] struct {
	m    int
	root *bpnode[K, V]
}

func (b *bptree[K, V]) Insert(k K, v V) bool {
	if b.root == nil {
		b.root = &bpnode[K, V]{leaf: true, keys: []K{k}, values: []V{v}}
		return false
	}
	// 找到对应范围的叶子节点
	node := b.root
	for {
		if !node.leaf {
			index := binsearch(node.keys, k)
			if index < len(node.keys) && node.keys[index] == k {
				node = node.children[index+1]
			} else {
				node = node.children[index]
			}
		} else {
			break
		}
	}

	// 找到对应的 key, 直接进行修改, 并返回
	index := binsearch(node.keys, k)
	if index < len(node.keys) && node.keys[index] == k {
		node.values[index] = v
		return true
	}

	// 分裂
	node.keys = insertValue(node.keys, index, k)
	node.values = insertValue(node.values, index, v)

	// 第二个判断条件是 root 节点的分列
	// 感觉根节点还是达到 m 之后再分裂比较好, 至少不会出现只有一个元素的节点.
	if len(node.keys) >= b.m {
		// 叶子节点分列
		leafLen := len(node.keys)
		half := leafLen / 2
		newNode := &bpnode[K, V]{leaf: true, keys: make([]K, leafLen-half), values: make([]V, leafLen-half), parent: node.parent}
		copy(newNode.keys, node.keys[half:])
		copy(newNode.values, node.values[half:])
		node.keys = node.keys[:half]
		node.values = node.values[:half]

		for {
			parent := node.parent
			if parent == nil {
				// 建立新的根节点
				newRoot := &bpnode[K, V]{leaf: false, keys: []K{newNode.keys[0]}, children: []*bpnode[K, V]{node, newNode}}
				b.root = newRoot
				node.parent = newRoot
				newNode.parent = newRoot
				break
			} else {
				// 索引节点插入新的节点
				index = binsearch(parent.keys, newNode.keys[0])
				parent.keys = insertValue(parent.keys, index, newNode.keys[0])
				parent.children = insertValue(parent.children, index+1, newNode)

				if len(parent.keys) < b.m-1 && parent.parent != nil {
					break
				}

				// 索引节点分裂
				node = parent
				notLeafLen := len(node.keys)
				half := len(node.keys) / 2

				// 始终保证 len(keys)+1 = len(children)
				newNode = &bpnode[K, V]{leaf: false, keys: make([]K, notLeafLen-half),
					children: make([]*bpnode[K, V], notLeafLen-half+1), parent: node.parent}
				copy(newNode.keys, node.keys[half:])
				// 新分裂出的节点, 最左侧是一个空节点.
				copy(newNode.children[1:], node.children[half+1:])
				node.keys = node.keys[:half]
				node.children = node.children[:half+1]
				newNode.children[0] = &bpnode[K, V]{leaf: newNode.children[1].leaf}
				// 修改索引
				for i := range newNode.children {
					newNode.children[i].parent = newNode
				}
				// 看到一个文档将空节点和兄弟节点进行关联, 形成链表可以解决空节点的问题.
				// 经过测试, 不管是采用 [a,b) 还是 (a,b] 那种方案, 按照我现在的实现都会产生空节点, 只是在左边还是右边的问题.
				// 使用链表来关联, 可以更好的处理这个问题, 但是我现在的采用的 (a,b] 感觉意义不大, 之后理解删除操作后在尝试看看如何修改.
				// newNode.children[0] = node
			}
		}
	}
	return false
}

func binsearch[K cmp.Ordered](values []K, k K) int {
	if values == nil {
		return 0
	}
	return sort.Search(len(values), func(i int) bool {
		return values[i] >= k
	})
}

func insertValue[V any](values []V, index int, v V) []V {
	values = append(values, v)
	if index == len(values)-1 {
		return values
	}
	copy(values[index+1:], values[index:])
	values[index] = v
	return values
}

func deleteValue[V any](values []V, index int) []V {
	return append(values[:index], values[index+1:]...)
}

func (b *bptree[K, V]) Search(k K) (V, bool) {
	var t V
	if b.root == nil {
		return t, false
	}
	node := b.root
	for {
		if node.leaf {
			index := binsearch(node.keys, k)
			if index < len(node.keys) && node.keys[index] == k {
				return node.values[index], true
			}
			return t, false
		} else {
			index := binsearch(node.keys, k)
			if index < len(node.keys) && node.keys[index] == k {
				node = node.children[index+1]
			} else {
				node = node.children[index]
			}
		}
	}
	return t, false
}

func (b *bptree[K, V]) borrow(l, n, r *bpnode[K, V]) *bpnode[K, V] {
	llen, rlen := l.keyLen(), r.keyLen()
	srcKey := n.keys[0]
	targetKey := n.keys[0]
	if rlen != 0 && rlen >= llen {
		srcKey = r.keys[0]
		if n.leaf {
			n.keys = append(n.keys, r.keys[0])
			r.keys = deleteValue(r.keys, 0)
			n.values = append(n.values, r.values[0])
			r.values = deleteValue(r.values, 0)
		} else {
			if r.children[0].keyLen() == 0 {
				n.keys = append(n.keys, r.keys[0])

				n.children = append(n.children, r.children[1])
				r.children[1].parent = n

				r.keys = deleteValue(r.keys, 0)
				r.children = deleteValue(r.children, 1)
			} else {
				n.children = append(n.children, r.children[0])
				r.children[0].parent = n
				n.keys = append(n.keys, r.children[0].keys[0])
				r.children[0] = &bpnode[K, V]{leaf: n.children[0].leaf, parent: r}
			}
		}
		targetKey = r.keys[0]
	} else if llen != 0 {
		srcKey = n.keys[0]
		if n.leaf {
			n.keys = insertValue(n.keys, 0, l.keys[len(l.keys)-1])
			n.values = insertValue(n.values, 0, l.values[len(l.values)-1])
			l.keys = deleteValue(l.keys, len(l.keys)-1)
			l.values = deleteValue(l.values, len(l.values)-1)
		} else {
			if n.children[0].keyLen() == 0 {
				n.keys = insertValue(n.keys, 0, l.keys[len(l.keys)-1])
				n.children = insertValue(n.children, 1, l.children[len(l.children)-1])
				l.children[len(l.children)-1].parent = n
				l.keys = deleteValue(l.keys, len(l.keys)-1)
				l.children = deleteValue(l.children, len(l.children)-1)
			} else {
				n.keys = insertValue(n.keys, 0, n.children[0].keys[0])
				n.children = insertValue(n.children, 0, &bpnode[K, V]{leaf: n.children[0].leaf, parent: n})
			}
		}
		targetKey = n.keys[0]
	} else {
		return n
	}
	if n.parent != nil {
		if n.parent != nil {
			index := binsearch(n.parent.keys, srcKey)
			if index == len(n.parent.keys) {
				n.parent.keys[index-1] = targetKey
			} else {
				if n.parent.keys[index] == srcKey {
					index++
				}
				n.parent.keys[index-1] = targetKey
			}
		}
	}
	return n
}

func (b *bptree[K, V]) merge(l, r *bpnode[K, V]) *bpnode[K, V] {

	// 合并节点内容
	if l.leaf {
		// 叶子节点可以直接合并
		l.keys = append(l.keys, r.keys...)
		l.values = append(l.values, r.values...)
	} else {
		if r.children[0].keyLen() == 0 {
			// 第一个子节点没有元素直接合并
			l.keys = append(l.keys, r.keys...)
			l.children = append(l.children, r.children[1:]...)
		} else {
			// 需要补充一个新的元素
			l.keys = append(l.keys, r.children[0].keys[0])
			l.keys = append(l.keys, r.keys...)
			l.children = append(l.children, r.children...)
		}

		for i := range l.children {
			l.children[i].parent = l
		}
	}
	// 调整父节点元素
	if l.parent == nil {
		return l
	}

	rFirstKey := r.keys[0]
	index := binsearch(l.parent.keys, rFirstKey)
	if index == len(l.parent.keys) {
		l.parent.keys = deleteValue(l.parent.keys, index-1)
		l.parent.children = deleteValue(l.parent.children, index)
	} else {
		if l.parent.keys[index] == rFirstKey {
			index++
		}
		l.parent.keys = deleteValue(l.parent.keys, index-1)
		l.parent.children = deleteValue(l.parent.children, index)
	}
	if l.parent.keyLen() == 0 {
		l.parent.keys = append(l.parent.keys, l.keys[0])
		l.parent.children = insertValue(l.parent.children, 0, &bpnode[K, V]{leaf: l.leaf, parent: l.parent})
	}
	return l
}

var tk = int64(110000)

func (b *bptree[K, V]) Delete(k K) bool {
	if b.root == nil {
		return false
	}
	//}
	// 找到对应范围的叶子节点
	node := b.root
	for {
		if !node.leaf {
			index := binsearch(node.keys, k)
			if index < len(node.keys) && node.keys[index] == k {
				node = node.children[index+1]
			} else {
				node = node.children[index]
			}
		} else {
			break
		}
	}

	// 找到对应的 key, 直接进行修改, 并返回
	index := binsearch(node.keys, k)
	if index < len(node.keys) && node.keys[index] == k {
		node.keys = deleteValue(node.keys, index)
		node.values = deleteValue(node.values, index)
	} else {
		return false
	}

	if node.parent == nil {
		return true
	}

	half := (b.m + 1) / 2

	if node.keyLen() < half {
		if node.keyLen() == 0 {
			index := binsearch(node.parent.keys, k)
			if index == len(node.parent.keys) {
				node.parent.keys = deleteValue(node.parent.keys, index-1)
				node.parent.children = deleteValue(node.parent.children, index)
			} else {
				if node.parent.keys[index] == k {
					index++
				}
				if index != 0 {
					node.parent.keys = deleteValue(node.parent.keys, index-1)
					node.parent.children = deleteValue(node.parent.children, index)
				}
			}
		} else {
			left, right := b.nearSibling(node)
			leftLen, nodeLen, rightLen := left.keyLen(), node.keyLen(), right.keyLen()
			// 从节点多的兄弟节点借一个元素
			// 我好想理解错误, 我把兄弟节点理解成包含堂兄弟节点了.
			if leftLen != 0 && leftLen+nodeLen <= b.m {
				node = b.merge(left, node)
			} else if rightLen != 0 && rightLen+nodeLen <= b.m {
				node = b.merge(node, right)
			} else {
				node = b.borrow(left, node, right)
			}
		}

		for {
			node = node.parent
			if node.parent == nil {
				if node.children[0].keyLen() == 0 {
					if len(node.children) == 2 {
						nr := node.children[1]
						nr.parent = nil
						b.root = nr
					} else {
						node.keys = node.keys[1:]
						node.children = node.children[1:]
					}
				}
				break
			}

			if node.keyLen() == 0 {
				index := binsearch(node.parent.keys, k)
				if index == len(node.parent.keys) {
					node.parent.keys = deleteValue(node.parent.keys, index-1)
					node.parent.children = deleteValue(node.parent.children, index)
				} else {
					if node.parent.keys[index] == k {
						index++
					}
					if index != 0 {
						node.parent.keys = deleteValue(node.parent.keys, index-1)
						node.parent.children = deleteValue(node.parent.children, index)
					}
				}
				continue
			}

			if node.keyLen() < half-1 {
				left, right := b.nearSibling(node)
				leftLen, nodeLen, rightLen := left.keyLen(), node.keyLen(), right.keyLen()
				if leftLen != 0 && leftLen+nodeLen <= b.m-1 {
					node = b.merge(left, node)
				} else if rightLen != 0 && rightLen+nodeLen <= b.m-1 {
					node = b.merge(node, right)
				} else {
					node = b.borrow(left, node, right)
				}
			} else {
				break
			}
		}

		if !b.root.leaf {
			if b.root.children[0].keyLen() == 0 {
				if len(b.root.children) == 2 {
					nr := b.root.children[1]
					nr.parent = nil
					b.root = nr
				} else {
					b.root.keys = b.root.keys[1:]
					b.root.children = b.root.children[1:]
				}
			} else if b.root.keyLen() == 0 {
				nr := b.root.children[0]
				nr.parent = nil
				b.root = nr
			}
		}

	}

	return true
}

func (b *bpnode[K, V]) keyLen() int {
	if b == nil {
		return 0
	}
	return len(b.keys)
}

func (b *bptree[K, V]) nearSibling(n *bpnode[K, V]) (l, r *bpnode[K, V]) {
	if n.parent == nil {
		return nil, nil
	}
	index := binsearch(n.parent.keys, n.keys[0])
	if index < n.parent.keyLen() && n.parent.keys[index] == n.keys[0] {
		index++
	}
	if index == len(n.parent.keys) {
		if n.parent.children[index-1].keyLen() == 0 {
			return nil, nil
		}
		return n.parent.children[index-1], nil
	}
	if index == 0 {
		return nil, n.parent.children[index+1]
	}

	if n.parent.children[index-1].keyLen() == 0 {
		return nil, n.parent.children[index+1]
	}
	return n.parent.children[index-1], n.parent.children[index+1]
}

func (b *bptree[K, V]) Iterator(f func(K, V) error) error {
	var dfs func(node *bpnode[K, V]) error
	dfs = func(node *bpnode[K, V]) error {
		if node.leaf {
			for i, k := range node.keys {
				if err := f(k, node.values[i]); err != nil {
					return err
				}
			}
		} else {
			for _, n := range node.children {
				if err := dfs(n); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return dfs(b.root)
}

func (b *bptree[K, V]) findNode(k K) *bpnode[K, V] {
	node := b.root
	for {
		if !node.leaf {
			index := binsearch(node.keys, k)
			if index < len(node.keys) && node.keys[index] == k {
				node = node.children[index+1]
			} else {
				node = node.children[index]
			}
		} else {
			break
		}
	}
	return node
}

func treePrint[K cmp.Ordered, V any](tree *bptree[K, V]) {
	if tree == nil {
		return
	}
	visited := map[string]bool{}
	var dfs func(node *bpnode[K, V], level int)
	dfs = func(node *bpnode[K, V], level int) {
		if visited[fmt.Sprintf("%p", node)] {
			fmt.Println(strings.Repeat("  ", level), node.keys)
			return
		}
		visited[fmt.Sprintf("%p", node)] = true
		if node.leaf {
			//fmt.Printf("%s%p,%p,%v\n", strings.Repeat("  ", level), node, node.parent, node.keys)
			fmt.Println(strings.Repeat("  ", level), node.keys)
		} else {
			//fmt.Printf("%s%p,%v,%p,%v\n", strings.Repeat("  ", level), node, node.keys, node.parent, node.children)
			fmt.Println(strings.Repeat("  ", level), node.keys)
			for _, n := range node.children {
				n := n
				dfs(n, level+1)
			}
		}
	}
	dfs(tree.root, 0)
}

func (b *bptree[K, V]) walk(fn func(n *bpnode[K, V]) error) error {
	var dfs func(node *bpnode[K, V]) error
	dfs = func(node *bpnode[K, V]) error {
		for _, n := range node.children {
			if err := dfs(n); err != nil {
				return err
			}
		}
		return fn(node)
	}
	return dfs(b.root)
}
