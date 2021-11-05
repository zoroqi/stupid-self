package q100

import (
	. "github.com/zoroqi/stupid-self"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	AssertEqual(t, SumNumbers(NewTreeNode([]int{4, 9, 0, 5, 1}, -1)), 1026)
	AssertEqual(t, SumNumbers(NewTreeNode([]int{4, 9, 0, 5}, -1)), 535)
	AssertEqual(t, SumNumbers(NewTreeNode([]int{4, 9, 0, 5, 1, -1, 6}, -1)), 1392)
	AssertEqual(t, SumNumbers(NewTreeNode([]int{4, 9, 0}, -1)), 89)
	AssertEqual(t, SumNumbers(NewTreeNode([]int{4, 9, 0, -1, 1, -1, 6}, -1)), 897)
	AssertEqual(t, SumNumbers(NewTreeNode([]int{4}, -1)), 4)
	AssertEqual(t, SumNumbers(NewTreeNode([]int{-1}, -1)), 0)
}
