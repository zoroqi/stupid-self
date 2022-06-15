package q800

import (
	. "github.com/zoroqi/stupid-self/golang"
	"reflect"
	"testing"
)

func TestPruneTree(t *testing.T) {
	f := reflect.DeepEqual(pruneTree(NewTreeNode([]int{1, 0, 1, 0, 0, 0, 1}, -1)), NewTreeNode([]int{1, -1, 1, -1, -1, -1, 1}, -1))
	AssertEqual(t, f, true)
}
