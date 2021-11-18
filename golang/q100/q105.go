package q100

import (
	. "github.com/zoroqi/stupid-self/golang"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	rv := preorder[0]
	var inRootIndex int
	for i, v := range inorder {
		if rv == v {
			inRootIndex = i
			break
		}
	}
	// leftLength := inRootIndex
	// rightLength := len(inorder) - inRootIndex - 1
	root.Left = BuildTree(preorder[1:inRootIndex+1], inorder[0:inRootIndex])
	root.Right = BuildTree(preorder[inRootIndex+1:], inorder[inRootIndex+1:])
	return root
}

func PreAndPostBuildTree(preorder, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	rv := preorder[1]
	var index int
	for i, v := range postorder {
		if rv == v {
			index = i
			break
		}
	}

	root.Left = PreAndPostBuildTree(preorder[1:index+2], postorder[0:index+1])
	// 最后一个元素在后续中不会使用到
	root.Right = PreAndPostBuildTree(preorder[index+2:], postorder[index+1:])
	return root
}
