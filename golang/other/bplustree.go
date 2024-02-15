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
	if len(node.keys) >= b.m || (node.leaf && node.parent == nil) {
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

				//if len(newNode.keys) > 1 {
				//	newNode.keys = newNode.keys[1:]
				//	newNode.children = newNode.children[1:]
				//} else {
				newNode.children[0] = &bpnode[K, V]{leaf: false}
				//}

				// 修改索引
				for i := range newNode.children {
					newNode.children[i].parent = newNode
				}

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
	if b.root == nil {
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
		node.keys = deleteValue(node.keys, index)
		node.children = deleteValue(node.children, index)
	} else {
		return false
	}

	return true
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

func treePrint[K cmp.Ordered, V any](tree *bptree[K, V]) {
	if tree == nil {
		return
	}
	var dfs func(node *bpnode[K, V], level int)
	dfs = func(node *bpnode[K, V], level int) {
		if node.leaf {
			//fmt.Printf("%s%p,%p,%v\n", strings.Repeat("  ", level), node, node.parent, node.keys)
			fmt.Println(strings.Repeat("  ", level), node.keys)
		} else {
			//fmt.Printf("%s%p,%v,%p,%v\n", strings.Repeat("  ", level), node, node.keys, node.parent, node.children)
			fmt.Println(strings.Repeat("  ", level), node.keys)
			for _, n := range node.children {
				dfs(n, level+1)
			}
		}
	}
	dfs(tree.root, 0)
}
