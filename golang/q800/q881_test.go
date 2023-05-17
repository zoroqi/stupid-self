package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestNumRescueBoats(t *testing.T) {
	stupid_self.AssertEqual(t, numRescueBoatsPlanA([]int{1, 2}, 3), 1)
	stupid_self.AssertEqual(t, numRescueBoatsPlanA([]int{3, 2, 2, 1}, 3), 3)
	stupid_self.AssertEqual(t, numRescueBoatsPlanA([]int{3, 5, 3, 4}, 5), 4)
	stupid_self.AssertEqual(t, numRescueBoatsPlanA([]int{1}, 5), 1)
	stupid_self.AssertEqual(t, numRescueBoatsPlanA([]int{1, 2, 3, 4, 5}, 5), 3)
}
