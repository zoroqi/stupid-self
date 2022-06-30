package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestCanVisitAllRoomsPlanA(t *testing.T) {
	stupid_self.AssertEqual(t, canVisitAllRoomsPlanB([][]int{{1}, {2}, {3}, {}}), true)
	stupid_self.AssertEqual(t, canVisitAllRoomsPlanB([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}), false)
}
