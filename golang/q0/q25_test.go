package q0

import (
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	list := stupid_self.NewListNode([]int{1, 2, 3, 4, 5, 6, 7,8})
	l := ReverseKGroup(list, 2)
	stupid_self.PrintListNode(l)
}
