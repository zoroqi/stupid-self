package q500

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestFindMode_Simple(t *testing.T) {
	stupid_self.AssertEqual(t,
		FindMode_Simple(stupid_self.NewTreeNode([]int{1, 0, 2, 0, 0, 2}, 0)),
		[]int{2})
}

func TestFindMode(t *testing.T) {
	stupid_self.AssertEqual(t,
		FindMode(stupid_self.NewTreeNode([]int{2, 1}, 0)),
		[]int{2, 1})
}
