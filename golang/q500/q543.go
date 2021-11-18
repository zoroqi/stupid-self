package q500

import . "github.com/zoroqi/stupid-self/golang"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, max := dep(root)
	// 这里max是最大节点数量, 不是长度, 需要减一
	return max - 1
}

func dep(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	ld, lm := dep(root.Left)
	rd, rm := dep(root.Right)

	m := MaxInt(lm, rm)
	nm := ld + rd + 1
	d := MaxInt(ld, rd) + 1
	if m > nm {
		return d, m
	}
	return d, nm
}
