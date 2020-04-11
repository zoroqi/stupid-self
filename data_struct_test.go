package stupid_self

import (
	"testing"
)

func TestBuildTreeNode(t *testing.T) {
	a := []int{1, 2, 2, 3, 4, 4, 3}
	n := NewTreeNode(a, 0)
	PrintTreeNode(n)
}

func TestListNode(t *testing.T) {
	array := []int{1,2,3,4,5,6,7}
	l := NewListNode(array)
	PrintListNode(l)
	l = ReverseListNode(l)
	PrintListNode(l)
}
