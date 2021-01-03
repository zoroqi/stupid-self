package q300

import (
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestMinPatches(t *testing.T) {
	stupid_self.AssertEqual(t, MinPatches([]int{1, 3, 7}, 20), 2)
	stupid_self.AssertEqual(t, MinPatches([]int{1, 5, 10}, 20), 2)
	stupid_self.AssertEqual(t, MinPatches([]int{1, 2, 2}, 5), 0)
	stupid_self.AssertEqual(t, MinPatches([]int{1, 2, 30}, 100), 4)
	stupid_self.AssertEqual(t, MinPatches([]int{}, 2), 2)
	stupid_self.AssertEqual(t, MinPatches([]int{}, 1025), 11)
	stupid_self.AssertEqual(t, MinPatches([]int{2}, 2147483647), 30)
	stupid_self.AssertEqual(t, MinPatches([]int{10}, 20), 4)
	stupid_self.AssertEqual(t, MinPatches([]int{1, 2, 2, 6, 34, 38, 41, 44, 47, 47, 56, 59, 62, 73, 77, 83, 87, 89, 94}, 20), 1)
}
