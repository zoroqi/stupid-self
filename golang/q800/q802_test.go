package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestEventualSafeNodesPlanA(t *testing.T) {
	stupid_self.AssertEqualFunc(t, eventualSafeNodesPlanA([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}),
		[]int{2, 4, 5, 6}, stupid_self.SetEqual)
	stupid_self.AssertEqualFunc(t, eventualSafeNodesPlanA([][]int{{}, {0, 2, 3, 4}, {3}, {4}, {}}),
		[]int{0, 1, 2, 3, 4}, stupid_self.SetEqual)
}
