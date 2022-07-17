package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestIsNStraightHand(t *testing.T) {
	stupid_self.AssertEqual(t, isNStraightHandPlanA([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3), true)
}
