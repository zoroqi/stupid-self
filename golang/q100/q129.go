package q100

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func SumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return innerSum(root, 0, 0)
}

func innerSum(node *TreeNode, sum, before int) int {
	current := before*10 + node.Val

	if node.Right == nil && node.Left == nil {
		return sum + current
	}
	if node.Left != nil {
		sum = innerSum(node.Left, sum, current)
	}
	if node.Right != nil {
		sum = innerSum(node.Right, sum, current)
	}
	return sum
}
