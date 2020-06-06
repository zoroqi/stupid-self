package q0

import (
	. "github.com/zoroqi/stupid-self"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	tmpHead := &ListNode{Val: 0, Next: head}
	tmp := tmpHead
	count := 0
	before := tmpHead
	var after *ListNode
	for tmp != nil && tmp.Next != nil {
		tmp = tmp.Next
		after = tmp.Next
		count++
		if count == k {
			count = 0
			s, e := reverseListNode(before.Next, k)
			before.Next = s
			e.Next = after
			before, tmp = e, e
		}
	}
	return tmpHead.Next
}

func reverseListNode(node *ListNode, k int) (*ListNode, *ListNode) {
	if k == 2 {
		t := node.Next
		t.Next = node
		node.Next = nil
		return t, t.Next
	}
	tmp1 := node
	tmp2 := node.Next
	i := 0
	for tmp2.Next != nil && i < k-2 {
		i++
		tmp1, tmp2, tmp2.Next = tmp2, tmp2.Next, tmp1
	}
	tmp2.Next = tmp1
	node.Next = nil
	return tmp2, node
}
