package q500

import (
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	stupid_self.AssertEqual(t,
		DiameterOfBinaryTree(stupid_self.NewTreeNode([]int{1, 2, 3, 4, 5}, 0)),
		3)
	stupid_self.AssertEqual(t,
		DiameterOfBinaryTree(stupid_self.NewTreeNode(
			[]int{1, 2, 3, 4, 5, 0, 0, 6, 0, 0, 7, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 9}, 0)),
		6)
}
