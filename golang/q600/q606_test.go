package q600

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func Test_tree2str(t *testing.T) {
	stupid_self.AssertEqual(t, tree2str(stupid_self.NewTreeNode([]int{1, 2, 3, 4}, 0)), "1(2(4))(3)")
	stupid_self.AssertEqual(t, tree2str(stupid_self.NewTreeNode([]int{1, 2, 3, 0, 4}, 0)), "1(2()(4))(3)")
}
