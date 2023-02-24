package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	stupid_self.AssertEqual(t, peakIndexInMountainArrayPlanA([]int{0, 2, 1, 0}), 1)
	stupid_self.AssertEqual(t, peakIndexInMountainArrayPlanB([]int{0, 2, 1, 0}), 1)
	stupid_self.AssertEqual(t, peakIndexInMountainArrayPlanC([]int{0, 2, 1, 0}), 1)
}
