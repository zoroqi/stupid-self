package q800

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func distanceKPlanB(root *TreeNode, target *TreeNode, k int) []int {
	if root == nil {
		return []int{}
	}
	type pathnode struct {
		n  *TreeNode // 节点
		d  int       // 距离 target 的距离
		lr int       // target 在左右子树, -1(左), 1(右), 0(target)
	}
	pns := []pathnode{}

	var find func(n *TreeNode, depth int) int
	find = func(n *TreeNode, depth int) int {
		if n == nil {
			return -1
		}
		if n.Val == target.Val {
			pns = append(pns, pathnode{n, 0, 0})
			return depth
		}
		lf := find(n.Left, depth+1)
		if lf >= 0 {
			pns = append(pns, pathnode{n, lf - depth, -1})
			return lf
		}
		rf := find(n.Right, depth+1)
		if rf >= 0 {
			pns = append(pns, pathnode{n, rf - depth, 1})
			return rf
		}
		return -1
	}

	find(root, 0)

	r := []int{}
	var dfs func(n *TreeNode, i int)
	dfs = func(n *TreeNode, i int) {
		if n == nil || i < 0 {
			return
		}
		if i == 0 {
			r = append(r, n.Val)
			return
		}
		dfs(n.Left, i-1)
		dfs(n.Right, i-1)
	}

	for _, v := range pns {
		// 计算根节点是否在其中
		if k-v.d == 0 {
			r = append(r, v.n.Val)
		}
		switch v.lr {
		case -1:
			dfs(v.n.Right, k-v.d-1)
		case 0:
			dfs(v.n.Left, k-v.d-1)
			dfs(v.n.Right, k-v.d-1)
		case 1:
			dfs(v.n.Left, k-v.d-1)
		}
	}

	return r
}
