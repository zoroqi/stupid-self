package q800

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func increasingBST(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	arr := []int{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	tmp := &TreeNode{}
	head := tmp
	for _, v := range arr {
		tmp.Right = &TreeNode{Val: v}
		tmp = tmp.Right
	}
	return head.Right
}

func increasingBSTPlanA(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (min, max *TreeNode)
	dfs = func(node *TreeNode) (min, max *TreeNode) {
		if node == nil {
			return nil, nil
		}
		if node.Left == nil && node.Right == nil {
			return node, node
		}
		lmin, lmax := dfs(node.Left)
		rmin, rmax := dfs(node.Right)

		node.Left = nil
		node.Right = rmin

		if lmax != nil {
			lmax.Right = node
		} else {
			lmin = node
		}
		if rmax == nil {
			rmax = node
		}
		return lmin, rmax
	}
	min, _ := dfs(root)
	return min
}
