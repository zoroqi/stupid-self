package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	stupid_self.AssertEqual(t,
		stupid_self.ListNodeToArray(middleNode(stupid_self.NewListNode([]int{1, 2, 3, 4, 5}))),
		[]int{3, 4, 5})
	stupid_self.AssertEqual(t,
		stupid_self.ListNodeToArray(middleNode(stupid_self.NewListNode([]int{1, 2, 3, 4, 5, 6}))),
		[]int{4, 5, 6})
}
