package q100

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	halfIndex, halfBefore := halfNode(head)
	halfBefore.Next = nil
	halfIndex = ReverseListNode(halfIndex)
	temp := &ListNode{Next: nil}
	newHead := temp
	tHead := head
	for tHead != nil && halfIndex != nil {
		temp.Next = tHead
		tHead = tHead.Next
		temp = temp.Next

		temp.Next = halfIndex
		halfIndex = halfIndex.Next
		temp = temp.Next
	}
	if halfIndex != nil {
		temp.Next = halfIndex
	} else {
		temp.Next = halfIndex
	}
	newHead.Next = nil
	newHead = nil
}

func halfNode(head *ListNode) (*ListNode, *ListNode) {
	endNode, halfNode := head, head
	halfBeforeNode := halfNode
	for endNode != nil && endNode.Next != nil {
		endNode = endNode.Next.Next
		halfBeforeNode = halfNode
		halfNode = halfNode.Next
	}
	return halfNode, halfBeforeNode
}
