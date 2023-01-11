package q800

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) <= 1 {
		return &TreeNode{Val: preorder[0]}
	}

	preMap := map[int]int{}
	postMap := map[int]int{}

	for i, v := range preorder {
		preMap[v] = i
	}
	for i, v := range postorder {
		postMap[v] = i
	}

	var rec func(preS, preE int, postS, postE int) *TreeNode
	rec = func(preS, preE int, postS, postE int) *TreeNode {
		if preS == preE {
			return &TreeNode{Val: preorder[preS]}
		}
		leftNodeS := preS + 1
		rightNodeE := postE - 1
		root := &TreeNode{Val: preorder[preS]}

		if preorder[leftNodeS] == postorder[rightNodeE] {
			root.Left = rec(leftNodeS, preE, postS, rightNodeE)
			return root
		}

		leftEnd := preMap[postorder[rightNodeE]]
		rightStart := postMap[preorder[leftNodeS]]

		root.Left = rec(leftNodeS, leftEnd-1, postS, rightStart)
		root.Right = rec(leftEnd, preE, rightStart+1, rightNodeE)

		return root
	}

	return rec(0, len(preorder)-1, 0, len(preorder)-1)
}
