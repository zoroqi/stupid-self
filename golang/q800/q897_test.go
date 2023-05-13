package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	stupid_self.AssertEqualFunc(t,
		increasingBST(stupid_self.NewTreeNode([]int{4, 2, 5, 1, 3, 6, 7}, 0)),
		stupid_self.NewTreeNode([]int{1, 0, 2, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7}, 0), stupid_self.TreeNodeEqual)
}

func TestIncreasingBSTPlanA(t *testing.T) {
	data := [][]int{{5, 3, 6, 2, 4, 0, 8, 1, 0, 0, 0, 7, 9}, {5, 1, 7}}
	for _, d := range data {
		stupid_self.AssertEqualFunc(t,
			increasingBST(stupid_self.NewTreeNode(d, 0)),
			increasingBSTPlanA(stupid_self.NewTreeNode(d, 0)), stupid_self.TreeNodeEqual)
	}
}
