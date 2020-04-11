package stupid_self

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构建二叉树
func NewTreeNode(a []int, nilValue int) *TreeNode {
	ta := make([]*TreeNode, 0, len(a))

	for _, v := range a {
		if v != nilValue {
			ta = append(ta, &TreeNode{Val: v})
		} else {
			ta = append(ta, nil)
		}
	}

	l := len(a)
	for i := 0; i < len(a); i++ {
		if ta[i] == nil {
			continue
		}
		if ((i+1)*2 - 1) < l {
			ta[i].Left = ta[(i+1)*2-1]
		} else {
			ta[i].Left = nil
		}

		if ((i + 1) * 2) < l {
			ta[i].Right = ta[(i+1)*2]
		} else {
			ta[i].Right = nil
		}
	}

	return ta[0]
}

// 打印二叉树, 前序遍历
func PrintTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("parent:%p ,%+v\n", root, *root)
	PrintTreeNode(root.Left)
	PrintTreeNode(root.Right)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 生成链表
func NewListNode(array []int) *ListNode {
	t := &ListNode{Val: 0}
	h := t
	for _, v := range array {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	t = h.Next
	h.Next = nil
	return t
}

func ListNodeToArray(node *ListNode) []int {
	var arr []int
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	return arr
}

// 打印链表
func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, ",")
		node = node.Next
	}
	fmt.Println()
}

// 翻转链表
func ReverseListNode(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}
	tmp1 := node
	tmp2 := node.Next

	for tmp2.Next != nil {
		tmp1, tmp2, tmp2.Next = tmp2, tmp2.Next, tmp1
	}
	tmp2.Next = tmp1
	node.Next = nil
	return tmp2
}
