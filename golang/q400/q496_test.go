package q400

import (
	"fmt"
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestNextGreaterElementPlanB(t *testing.T) {
	fmt.Println(nextGreaterElementPlanB([]int{3, 1, 2, 5, 4, 6}, []int{3, 1, 2, 5, 4, 6}))
}

func TestNextGreaterElementPlanC(t *testing.T) {
	stupid_self.AssertEqual(t, []int{5, 5, 6, 6, -1}, nextGreaterElementPlanC([]int{3, 2, 5, 4, 6}, []int{3, 1, 2, 5, 4, 6}))
}
