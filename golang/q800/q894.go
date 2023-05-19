package q800

import (
	"fmt"
	. "github.com/zoroqi/stupid-self/golang"
	"sort"
)

func allPossibleFBTPlanA(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	three := &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}
	if n == 3 {
		return []*TreeNode{three}
	}
	var clone func(node *TreeNode) *TreeNode
	clone = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		return &TreeNode{
			Left:  clone(root.Left),
			Right: clone(root.Right),
		}
	}

	var addLeft func(root, node *TreeNode) []*TreeNode
	addLeft = func(root, node *TreeNode) (r []*TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			node.Left = &TreeNode{}
			node.Right = &TreeNode{}
			ntree := clone(root)
			node.Left = nil
			node.Right = nil
			return []*TreeNode{ntree}
		} else {
			nl := addLeft(root, node.Left)
			nr := addLeft(root, node.Right)
			if nl != nil {
				r = append(r, nl...)
			}
			if nr != nil {
				r = append(r, nr...)
			}
		}
		return
	}

	toString := func(node *TreeNode) string {
		return node.String()
	}

	dup := map[string]bool{}
	r := make([]*TreeNode, 0)
	r = append(r, three)
	for i := 5; i <= n; i += 2 {
		nr := make([]*TreeNode, 0, len(r)*2)
		for _, v := range r {
			rr := addLeft(v, v)
			for _, ntree := range rr {
				k := toString(ntree)
				if !dup[k] {
					nr = append(nr, ntree)
					dup[k] = true
				}

			}
		}
		r = nr
	}

	return r
}

func allPossibleFBTPlanB(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	dup := map[string]bool{}
	r := [][]int{{0}}
	for i := 3; i <= n; i += 2 {
		tmp := make([][]int, 0, len(r)+2)
		for _, v := range r {
			l := len(v) - 1
			for i := 0; i <= l; i++ {
				newA := make([]int, l+2)
				copy(newA, v[:i])
				copy(newA[i:], v[i+1:])
				newA[l] = (v[i]+1)*2 - 1
				newA[l+1] = (v[i] + 1) * 2
				sort.Ints(newA)
				k := fmt.Sprint(newA)
				if !dup[k] {
					tmp = append(tmp, newA)
					dup[k] = true
				}
			}
		}
		r = tmp
	}
	rr := make([]*TreeNode, len(r))
	buildTree := func(arr []int) *TreeNode {
		nodes := map[int]*TreeNode{}
		for _, index := range arr {
			n := index
			nodes[n] = &TreeNode{}
			for n != 0 {
				if n%2 == 0 {
					pi := n/2 - 1
					if pn := nodes[pi]; pn != nil {
						nodes[pi].Right = nodes[n]
						break
					} else {
						nodes[pi] = &TreeNode{}
						nodes[pi].Right = nodes[n]
						n = pi
					}
				} else {
					pi := (n+1)/2 - 1
					if pn := nodes[pi]; pn != nil {
						nodes[pi].Left = nodes[n]
						break
					} else {
						nodes[pi] = &TreeNode{}
						nodes[pi].Left = nodes[n]
						n = pi
					}
				}
			}
		}
		return nodes[0]
	}
	for i := 0; i < len(r); i++ {
		rr[i] = buildTree(r[i])
	}
	return rr
}

func allPossibleFBTPlanC(n int) []*TreeNode {
	f := make([][]*TreeNode, n+1)
	f[1] = []*TreeNode{&TreeNode{}}
	var dfs func(int) []*TreeNode
	dfs = func(num int) []*TreeNode {
		if f[num] != nil {
			return f[num]
		}
		result := []*TreeNode{}
		for i := 0; i < num-1; i++ {
			left := dfs(i)
			right := dfs(num - 1 - i)
			for _, lv := range left {
				for _, rv := range right {
					result = append(result, &TreeNode{Left: lv, Right: rv})
				}
			}
		}
		f[num] = result
		return result
	}
	return dfs(n)
}
