package q500

import (
	. "github.com/zoroqi/stupid-self/golang"
	"math"
)

func FindBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stacks := []*TreeNode{root}
	var node *TreeNode
	for len(stacks) > 0 {
		node = stacks[0]
		stacks = stacks[1:]
		if node.Right != nil {
			stacks = append(stacks, node.Right)
		}
		if node.Left != nil {
			stacks = append(stacks, node.Left)
		}
	}
	return node.Val
}

func FindBottomLeftValue_first(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var r int
	flag := false
	queue := make([]*TreeNode, 2)
	queue[0] = nil
	queue[1] = root
	for len(queue) > 1 {
		q := queue[0]
		queue = queue[1:]
		if q == nil {
			queue = append(queue, nil)
			flag = true
			continue
		}
		if flag {
			r = q.Val
			flag = false
		}
		if q.Left != nil {
			queue = append(queue, q.Left)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}

	return r
}

func LargestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var r []int
	max := math.MinInt32

	queue := make([]*TreeNode, 2)
	queue[0] = root
	queue[1] = nil
	for len(queue) > 1 {
		q := queue[0]
		queue = queue[1:]
		if q == nil {
			queue = append(queue, nil)
			r = append(r, max)
			max = math.MinInt32
			continue
		}
		if max < q.Val {
			max = q.Val
		}
		if q.Left != nil {
			queue = append(queue, q.Left)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}
	r = append(r, max)
	return r
}
