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
	if !(2 < m && m < 100) {
		m = 10
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
				// 新分裂出的节点, 最左侧是以个空节点.
				copy(newNode.children[1:], node.children[half+1:])
				node.keys = node.keys[:half]
				node.children = node.children[:half+1]

				newNode.children[0] = &bpnode[K, V]{leaf: node.leaf}
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

func (b *bptree[K, V]) Delete(k K) bool {
	//fmt.Println("+++++")
	fmt.Println("delete ", k)
	if b.root == nil {
		return false
	}
	// 找到对应范围的叶子节点
	node := b.root
	for {
		if !node.leaf {
			index := binsearch(node.keys, k)
			fmt.Println(node.keys, k, index)
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
	h2 := b.m / 2
	if len(node.keys) < h2 {
		left, right := b.nearSibling(node)
		leftLen, _, rightLen := left.keyLen(), node.keyLen(), right.keyLen()
		// 从节点多的兄弟节点借一个元素
		// 我好想理解错误, 我把兄弟节点理解成包含堂兄弟节点了.
		sibling := rightLen > leftLen // true: 从右节点借, false: 从左节点借
		//fmt.Println(left, right)
		if leftLen != 0 && leftLen <= h2 {
			left.keys = append(left.keys, node.keys...)
			left.values = append(left.values, node.values...)
			index := binsearch(left.parent.keys, node.keys[0])
			if index == len(left.parent.keys) {
				left.parent.keys = deleteValue(left.parent.keys, index-1)
				left.parent.children = deleteValue(left.parent.children, index)
			} else {
				if left.parent.keys[index] == node.keys[0] {
					index++
				}
				left.parent.keys = deleteValue(left.parent.keys, index-1)
				left.parent.children = deleteValue(left.parent.children, index)
			}
			node = left
		} else if rightLen != 0 && rightLen <= h2 {
			node.keys = append(node.keys, right.keys...)
			node.values = append(node.values, right.values...)
			//index := binsearch(node.parent.keys, right.keys[0])
			//node.parent.keys = deleteValue(node.parent.keys, index)
			//node.parent.children = deleteValue(node.parent.children, index+1)

			index := binsearch(node.parent.keys, right.keys[0])
			if index == len(node.parent.keys) {
				node.parent.keys = deleteValue(node.parent.keys, index-1)
				node.parent.children = deleteValue(node.parent.children, index)
			} else {
				if node.parent.keys[index] == right.keys[0] {
					index++
				}
				node.parent.keys = deleteValue(node.parent.keys, index-1)
				node.parent.children = deleteValue(node.parent.children, index)
			}
			if node.parent.keyLen() == 0 {
				node.parent.keys = append(node.parent.keys, node.keys[0])
				node.parent.children = insertValue(node.parent.children, 0, &bpnode[K, V]{leaf: node.leaf, parent: node.parent})
			}
		} else if sibling {
			rkey := right.keys[0]
			node.keys = append(node.keys, rkey)
			node.values = append(node.values, right.values[0])
			right.keys = deleteValue(right.keys, 0)
			right.values = deleteValue(right.values, 0)
			rindex := binsearch(right.parent.keys, rkey)
			if rindex == len(right.parent.keys) {
				right.parent.keys[rindex-1] = rkey
			} else {
				if right.parent.keys[rindex] == rkey {
					rindex++
				}
				right.parent.keys[rindex-1] = rkey
			}
		} else if leftLen > 0 {
			lkey := left.keys[len(left.keys)-1]
			node.keys = insertValue(node.keys, 0, lkey)
			node.values = insertValue(node.values, 0, left.values[len(left.values)-1])
			nindex := binsearch(node.parent.keys, node.keys[0])
			if nindex == len(node.parent.keys) {
				node.parent.keys[nindex-1] = node.keys[0]
			} else {
				if node.keys[0] == node.parent.keys[nindex] {
					nindex++
				}
				node.parent.keys[nindex] = node.keys[0]
			}

			left.keys = deleteValue(left.keys, len(left.keys)-1)
			left.values = deleteValue(left.values, len(left.keys)-1)
		}
		//fmt.Println(node.keys)
		for {
			treePrint(b)
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
			half := (b.m - 1) / 2
			if node.keyLen() <= half {
				left, right := b.nearSibling(node)
				leftLen, _, rightLen := left.keyLen(), node.keyLen(), right.keyLen()
				if leftLen == rightLen && leftLen == 0 {
					break
				}

				// 从节点多的兄弟节点借一个元素
				// 我好想理解错误, 我把兄弟节点理解成包含堂兄弟节点了.
				sibling := rightLen > leftLen // true: 从右节点借, false: 从左节点借
				//fmt.Println(sibling, node.parent.keys, node.keys, nodeLen, leftLen, rightLen, left, right)
				if leftLen != 0 && leftLen <= half {
					left.keys = append(left.keys, node.keys...)
					if node.children[0].keyLen() == 0 {
						left.children = append(left.children, node.children[1:]...)
					} else {
						left.children = append(left.children, node.children...)
					}
					for i := range left.children {
						left.children[i].parent = left
					}
					index := binsearch(left.parent.keys, node.keys[0])
					if index == len(left.parent.keys) {
						left.parent.keys = deleteValue(left.parent.keys, index-1)
						left.parent.children = deleteValue(left.parent.children, index)
					} else {
						if left.parent.keys[index] == node.keys[0] {
							index++
						}
						left.parent.keys = deleteValue(left.parent.keys, index-1)
						left.parent.children = deleteValue(left.parent.children, index)
					}
					node = left
				} else if rightLen != 0 && rightLen <= half {
					node.keys = append(node.keys, right.keys...)
					if right.children[0].keyLen() == 0 {
						node.children = append(node.children, right.children[1:]...)
					} else {
						node.children = append(node.children, right.children...)
					}
					for i := range node.children {
						node.children[i].parent = node
					}
					index := binsearch(node.parent.keys, right.keys[0])
					if index == len(node.parent.keys) {
						node.parent.keys = deleteValue(node.parent.keys, index-1)
						node.parent.children = deleteValue(node.parent.children, index)
					} else {
						if node.parent.keys[index] == right.keys[0] {
							index++
						}
						node.parent.keys = deleteValue(node.parent.keys, index-1)
						node.parent.children = deleteValue(node.parent.children, index)
					}
				} else if sibling {
					rkey := right.keys[0]
					node.keys = append(node.keys, rkey)
					if right.children[0].keyLen() == 0 {
						right.children[1].parent = node
						node.children = append(node.children, right.children[1])
						right.keys = deleteValue(right.keys, 0)
						right.children = deleteValue(right.children, 1)
					} else {
						right.children[0].parent = node
						node.children = append(node.children, right.children[0])
						right.keys = deleteValue(right.keys, 0)
						right.children = deleteValue(right.children, 0)
					}
					rindex := binsearch(right.parent.keys, rkey)
					right.parent.keys[rindex] = right.keys[0]
				} else {
					lkey := left.keys[len(left.keys)-1]
					node.keys = insertValue(node.keys, 0, lkey)
					if node.children[0].keyLen() == 0 {
						node.children = insertValue(node.children, 1, left.children[len(left.children)-1])
					} else {
						node.children = insertValue(node.children, 0, left.children[len(left.children)-1])
					}
					left.children[len(left.children)-1].parent = node
					nindex := binsearch(node.parent.keys, node.keys[0])
					node.parent.keys[nindex] = node.keys[0]
					left.keys = deleteValue(left.keys, len(left.keys)-1)
					left.children = deleteValue(left.children, len(left.children)-1)
				}
				if node.parent.keyLen() == 0 {
					node.parent.keys = append(node.parent.keys, node.keys[0])
					node.parent.children = insertValue(node.parent.children, 0, &bpnode[K, V]{leaf: node.leaf, parent: node.parent})
				}
			} else {
				break
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
	fmt.Println(n.parent.keys, n.keys)
	index := binsearch(n.parent.keys, n.keys[0])
	fmt.Println(index, n.keys[0], n.parent.keys)
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

// 索引节点直接构成链表更容易找到兄弟节点
func (b *bptree[K, V]) rightSibling(n *bpnode[K, V]) *bpnode[K, V] {
	node := n
	depth := 0
	up := true
	// 没有直接查到就向上, 之后在向下.
	// 向上的过程始终找右测, 向下直接去 children[0] 的元素(需要过滤掉空节点)
	for {
		if node.parent == nil {
			return nil
		}
		if up {
			index := binsearch(node.parent.keys, node.keys[0])
			if index < len(node.parent.keys) && node.parent.keys[index] >= node.keys[0] {
				index++
			}
			depth++
			if index+1 < len(node.parent.children) {
				node = node.parent.children[index+1]
				depth--
				if depth == 0 {
					return node
				}
				up = false
			} else {
				node = node.parent
			}
		} else {
			if depth == 0 {
				if len(node.keys) == 0 {
					return node.parent.children[1]
				}
				return node
			} else {
				if len(node.keys) == 0 {
					node = node.parent.children[1]
					node = node.children[0]
				} else {
					node = node.children[0]
				}
			}
			depth--
		}
	}

	return nil
}

func (b *bptree[K, V]) leftSibling(n *bpnode[K, V]) *bpnode[K, V] {
	node := n
	depth := 0
	up := true
	// 没有直接查到就向上, 之后在向下.
	// 向上的过程始终找右, 向下返回 children[n-1] 的元素
	for {
		if node.parent == nil {
			return nil
		}
		if up {
			index := binsearch(node.parent.keys, node.keys[0])
			depth++
			if index > 0 {
				node = node.parent.children[index]
				depth--
				if depth == 0 && len(node.keys) != 0 {
					return node
				}
				up = false
			} else {
				node = node.parent
			}
		} else {
			if depth == 0 {
				return node
			} else {
				node = node.children[len(node.children)-1]
			}
			depth--
		}
	}

	return nil
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
