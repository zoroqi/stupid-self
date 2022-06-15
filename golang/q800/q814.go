package q800

import . "github.com/zoroqi/stupid-self/golang"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)
		if root.Right == nil && root.Left == nil && root.Val == 0 {
			return nil
		}
		return root
	}
	return dfs(root)
}
