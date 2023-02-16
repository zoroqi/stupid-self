package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestLongestMountain(t *testing.T) {
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{2, 1, 4, 7, 3, 2, 5}), 5)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{2, 2, 2}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{3, 2, 1}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{1, 2, 3}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{2, 2, 3}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{2, 2, 1}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{3, 2, 2}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{1, 2, 2}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{2, 1, 3}), 0)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{1, 3, 2}), 3)
	stupid_self.AssertEqual(t, longestMountainPlanB([]int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}), 11)
}
