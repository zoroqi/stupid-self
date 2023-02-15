package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestName(t *testing.T) {
	stupid_self.AssertEqual(t, leafSimilar(
		stupid_self.NewTreeNode([]int{3, 5, 1, 6, 2, 9, 8, 0, 0, 7, 4}, 0),
		stupid_self.NewTreeNode([]int{3, 5, 1, 6, 7, 4, 2, 0, 0, 0, 0, 0, 0, 9, 8}, 0),
	), true)

	stupid_self.AssertEqual(t, leafSimilar(
		stupid_self.NewTreeNode([]int{1, 2, 3}, 0),
		stupid_self.NewTreeNode([]int{1, 3, 2}, 0),
	), false)
}
