package q100

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	root := stupid_self.NewTreeNode([]int{7, 3, 15, 0, 0, 9, 20}, 0)
	iter := Constructor(root)
	for iter.HasNext() {
		t.Log(iter.Next())
	}
}
