package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestIsMonotonic(t *testing.T) {
	stupid_self.AssertEqual(t, isMonotonic([]int{1, 2, 2, 3}), true)
	stupid_self.AssertEqual(t, isMonotonic([]int{6, 5, 4, 4}), true)
	stupid_self.AssertEqual(t, isMonotonic([]int{1, 3, 2}), false)
	stupid_self.AssertEqual(t, isMonotonic([]int{1, 1, 1, 3, 3, 3, 2}), false)
	stupid_self.AssertEqual(t, isMonotonic([]int{2, 2, 2, 3, 3, 4, 4}), true)
}
