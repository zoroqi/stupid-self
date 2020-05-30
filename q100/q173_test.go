package q100

import (
	"fmt"
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	root := stupid_self.NewTreeNode([]int{7,3,15,0,0,9,20},0)
	//stupid_self.InOrder(root)
	iter := Constructor(root)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
