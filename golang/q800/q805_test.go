package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestSplitArraySameAveragePlanA(t *testing.T) {
	stupid_self.AssertEqual(t, splitArraySameAveragePlanA([]int{1, 2, 3, 4, 5, 6, 7, 8}), true)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanA([]int{2, 0, 5, 6, 16, 12, 15, 12, 4}), true)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanA([]int{3, 1}), false)
}

func TestSplitArraySameAveragePlanB(t *testing.T) {
	stupid_self.AssertEqual(t, splitArraySameAveragePlanB([]int{1, 2, 3, 4, 5, 6, 7, 8}), true)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanB([]int{2, 0, 5, 6, 16, 12, 15, 12, 4}), true)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanB([]int{3, 1}), false)
}
func TestSplitArraySameAveragePlanC(t *testing.T) {
	stupid_self.AssertEqual(t, splitArraySameAveragePlanC([]int{1, 2, 3, 4, 5, 6, 7, 8}), true)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanC([]int{2, 0, 5, 6, 16, 12, 15, 12, 4}), true)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanC([]int{3, 1}), false)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanC([]int{60, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}), false)
	stupid_self.AssertEqual(t, splitArraySameAveragePlanC([]int{1, 2, 3, 4, 7, 8}), false)
}
