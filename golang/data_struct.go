package stupid_self

import (
	"errors"
	"fmt"
	"strings"
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
	fmt.Printf("node:%p ,%+v\n", root, *root)
	PrintTreeNode(root.Left)
	PrintTreeNode(root.Right)
}

func PrintDeepTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	var dfs func(int, *TreeNode)
	dfs = func(i int, node *TreeNode) {
		if node == nil {
			return
		}
		fmt.Printf("%2d%snode:%p ,%+v\n", i, strings.Repeat(" ", i), node, *node)
		dfs(i+1, node.Left)
		dfs(i+1, node.Right)
	}
	dfs(0, root)
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

func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func (t *TreeNode) String() string {
	var dfs func(node *TreeNode)
	sb := strings.Builder{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		sb.WriteString(fmt.Sprintf("%d ", node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(t)
	return sb.String()
}

func InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Println(node.Val)
	InOrder(node.Right)
}

func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Val)
}

type Stack[T any] struct {
	arr []T
}

var (
	StackEmptyError = errors.New("stack empty")
)

func (s *Stack[T]) Push(t T) int {
	s.arr = append(s.arr, t)
	return s.Length()
}

func (s *Stack[T]) Pop() (T, error) {
	l := len(s.arr)
	if l == 0 {
		var t T
		return t, StackEmptyError

	}
	t := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return t, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Length() int {
	return len(s.arr)
}

func (s *Stack[T]) GetPop() (T, error) {
	l := len(s.arr)
	if l == 0 {
		var t T
		return t, StackEmptyError

	}
	return s.arr[l-1], nil
}
func (s *Stack[T]) Print() {
	fmt.Println(s.arr)
}
