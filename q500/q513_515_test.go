package q500

import (
	. "github.com/zoroqi/stupid-self"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	AssertEqual(t, FindBottomLeftValue(NewTreeNode([]int{2, 1, 3}, 0)), 1)
	AssertEqual(t, FindBottomLeftValue(NewTreeNode([]int{1, 2, 3, 4, 0, 5, 6, 0, 8, 0, 0, 7}, 0)), 8)
	AssertEqual(t, FindBottomLeftValue(NewTreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)), 8)
	AssertEqual(t, FindBottomLeftValue(NewTreeNode([]int{1}, 0)), 1)
}

func TestLargestValues(t *testing.T) {
	AssertEqual(t, LargestValues(NewTreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)), []int{1, 3, 7, 9})
	AssertEqual(t, LargestValues(NewTreeNode([]int{1, 3, 2, 5, 3, 0, 9}, 0)), []int{1, 3, 9})
}
