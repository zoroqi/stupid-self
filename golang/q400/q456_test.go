package q400

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestFind132patternPlanA(t *testing.T) {
	stupid_self.AssertEqual(t, find132patternPlanA([]int{3, 1, 4, 2}), true)
	stupid_self.AssertEqual(t, find132patternPlanA([]int{-1, 3, 2, 0}), true)
	stupid_self.AssertEqual(t, find132patternPlanA([]int{1, 2, 3, 4}), false)
}
