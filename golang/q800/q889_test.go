package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestConstructFromPrePost(t *testing.T) {
	stupid_self.AssertEqualFunc(t,
		constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}),
		stupid_self.NewTreeNode([]int{1, 2, 3, 4, 5, 6, 7}, 0), stupid_self.TreeNodeEqual)
	stupid_self.AssertEqualFunc(t,
		constructFromPrePost([]int{1}, []int{1}),
		stupid_self.NewTreeNode([]int{1}, 0), stupid_self.TreeNodeEqual)
}
