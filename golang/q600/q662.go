package q600

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func widthOfBinaryTreePlanA(root *TreeNode) int {
	type item struct {
		index int
		node  *TreeNode
	}
	queue := []item{}
	queue = append(queue, item{
		index: 0,
		node:  root,
	}, item{index: -1})
	leftI := -1
	rightI := -1
	max := 0
	for len(queue) > 1 {
		n := queue[0]
		if n.index == -1 {
			l := rightI - leftI + 1
			if l > max {
				max = l
			}
			leftI = -1
			rightI = -1
			queue = append(queue[1:], n)
			continue
		}
		if leftI == -1 {
			leftI = n.index
		}
		rightI = n.index
		queue = queue[1:]
		if n.node.Left != nil {
			queue = append(queue, item{
				index: n.index * 2,
				node:  n.node.Left,
			})
		}
		if n.node.Right != nil {
			queue = append(queue, item{
				index: n.index*2 + 1,
				node:  n.node.Right,
			})
		}
	}
	l := rightI - leftI + 1
	if l > max {
		max = l
	}
	return max

}

func widthOfBinaryTreePlanB(root *TreeNode) int {
	type item struct {
		index int
		node  *TreeNode
	}
	queue := []item{}
	queue = append(queue, item{
		index: 0,
		node:  root,
	})
	max := 1
	for len(queue) != 0 {
		l := queue[len(queue)-1].index - queue[0].index + 1
		if l > max {
			max = l
		}
		q := []item{}
		for _, i := range queue {
			if i.node.Left != nil {
				q = append(q, item{
					index: i.index * 2, node: i.node.Left,
				})
			}
			if i.node.Right != nil {
				q = append(q, item{
					index: i.index*2 + 1, node: i.node.Right,
				})
			}
		}
		queue = q
	}
	return max
}
