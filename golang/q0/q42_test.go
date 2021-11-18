package q0

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestTrap(t *testing.T) {
	stupid_self.AssertEqual(t, Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
	stupid_self.AssertEqual(t, Trap([]int{0, 1, 0, 2, 1, 0, 3, 1, 3, 1, 2, 1}), 7)
	stupid_self.AssertEqual(t, Trap([]int{0, 1, 2, 3, 4}), 0)
	stupid_self.AssertEqual(t, Trap([]int{4, 1, 2, 3, 1, 4}), 9)
	stupid_self.AssertEqual(t, Trap([]int{4, 2, 3}), 1)
}
