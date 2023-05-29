package q1900

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestOperationsOnTree(t *testing.T) {
	tree := Q1933Constructor([]int{-1, 0, 0, 1, 1, 2, 2})
	stupid_self.AssertEqual(t, tree.Lock(2, 2), true)
	stupid_self.AssertEqual(t, tree.Lock(2, 2), false)
	stupid_self.AssertEqual(t, tree.Unlock(2, 2), true)
	stupid_self.AssertEqual(t, tree.Unlock(2, 2), false)
	stupid_self.AssertEqual(t, tree.Lock(2, 2), true)
	stupid_self.AssertEqual(t, tree.Upgrade(2, 2), false)
	stupid_self.AssertEqual(t, tree.Unlock(2, 2), true)

	stupid_self.AssertEqual(t, tree.Lock(2, 2), true)
	stupid_self.AssertEqual(t, tree.Unlock(2, 3), false)
	stupid_self.AssertEqual(t, tree.Unlock(2, 2), true)
	stupid_self.AssertEqual(t, tree.Lock(4, 5), true)
	stupid_self.AssertEqual(t, tree.Upgrade(0, 1), true)
	stupid_self.AssertEqual(t, tree.Lock(0, 1), false)
}
