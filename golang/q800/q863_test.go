package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestDistanceK(t *testing.T) {
	stupid_self.AssertEqualFunc(t,
		distanceKPlanB(stupid_self.NewTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, -1), &stupid_self.TreeNode{Val: 5}, 2),
		[]int{7, 4, 1},
		stupid_self.SetEqual)
	stupid_self.AssertEqualFunc(t,
		distanceKPlanB(stupid_self.NewTreeNode([]int{1}, -1), &stupid_self.TreeNode{Val: 1}, 3),
		[]int{},
		stupid_self.SetEqual)
	stupid_self.AssertEqualFunc(t,
		distanceKPlanB(stupid_self.NewTreeNode([]int{1}, -1), &stupid_self.TreeNode{Val: 1}, 0),
		[]int{1},
		stupid_self.SetEqual)
}
