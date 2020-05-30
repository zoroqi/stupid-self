package q100

import . "github.com/zoroqi/stupid-self"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	root   *TreeNode
	stack  []*TreeNode
	offset int
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return BSTIterator{root: root, stack: stack, offset: -1}
	}
	offset := -1
	for n := root; n != nil; n = n.Left {
		offset++
		stack = append(stack, n)
	}
	return BSTIterator{root: root, stack: stack, offset: offset}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack[this.offset]
	this.offset--
	for n := node.Right; n != nil; n = n.Left {
		this.offset++
		if this.offset < len(this.stack) {
			this.stack[this.offset] = n
		} else {
			this.stack = append(this.stack, n)
		}
	}

	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.offset == -1 {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
