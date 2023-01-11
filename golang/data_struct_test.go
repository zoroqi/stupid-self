package stupid_self

import (
	"testing"
)

func TestBuildTreeNode(t *testing.T) {
	a := []int{1, 0, 2, 3, 4, 4, 3}
	n := NewTreeNode(a, 0)
	PrintTreeNode(n)
}

func TestListNode(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	l := NewListNode(array)
	PrintListNode(l)
	l = ReverseListNode(l)
	PrintListNode(l)
}

func TestStack(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	p1, _ := stack.GetPop()
	AssertEqual(t, p1, 3)
	AssertEqual(t, stack.IsEmpty(), false)
	AssertEqual(t, stack.Length(), 3)
	p2, _ := stack.Pop()
	AssertEqual(t, p2, 3)
	AssertEqual(t, stack.IsEmpty(), false)
	AssertEqual(t, stack.Length(), 2)
	stack.Pop()
	stack.Pop()
	AssertEqual(t, stack.IsEmpty(), true)
	AssertEqual(t, stack.Length(), 0)
	_, err := stack.Pop()
	AssertEqual(t, err, StackEmptyError)
}
