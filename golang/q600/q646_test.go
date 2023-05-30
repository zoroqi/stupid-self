package q600

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestDFindLongestChain(t *testing.T) {
	stupid_self.AssertEqual(t, findLongestChainPlanC([][]int{{1, 2}, {2, 3}, {3, 4}}), 2)
	stupid_self.AssertEqual(t, findLongestChainPlanC([][]int{{1, 2}, {7, 8}, {4, 5}}), 3)
	stupid_self.AssertEqual(t, findLongestChainPlanC([][]int{{7, 9}, {4, 5}, {7, 9}, {-7, -1}, {0, 10}, {3, 10}, {3, 6}, {2, 3}}), 4)
	stupid_self.AssertEqual(t, findLongestChainPlanC([][]int{{1, 2}, {2, 4}, {2, 6}}), 1)
	stupid_self.AssertEqual(t, findLongestChainPlanC([][]int{{0, 0}, {0, 0}}), 1)
	stupid_self.AssertEqual(t, findLongestChainPlanC([][]int{{0, 0}, {1, 20}, {2, 3}, {4, 5}}), 3)
}
