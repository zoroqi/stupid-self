package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestMinSwap(t *testing.T) {
	stupid_self.AssertEqual(t, minSwap([]int{1, 2, 3, 8}, []int{5, 6, 7, 4}), 1)
	stupid_self.AssertEqual(t, minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}), 1)
	stupid_self.AssertEqual(t, minSwap([]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}), 1)
	stupid_self.AssertEqual(t, minSwap([]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}), 1)
	stupid_self.AssertEqual(t, minSwap([]int{0, 4, 4, 5, 9}, []int{0, 1, 6, 8, 10}), 1)
}

func TestMinSwapPlanA(t *testing.T) {
	stupid_self.AssertEqual(t, IsSequences(minSwapPlanA([]int{1, 4, 5}, []int{3, 2, 5})), true)
	stupid_self.AssertEqual(t, IsSequences(minSwapPlanA([]int{1, 4, 3}, []int{3, 2, 7})), true)
	stupid_self.AssertEqual(t, IsSequences(minSwapPlanA([]int{1, 3, 5, 4}, []int{1, 2, 3, 7})), true)
	stupid_self.AssertEqual(t, IsSequences(minSwapPlanA([]int{0, 4, 4, 5, 9}, []int{0, 1, 6, 8, 10})), true)
	stupid_self.AssertEqual(t, IsSequences(minSwapPlanA([]int{0, 7, 3, 4, 5}, []int{0, 2, 9, 10, 11})), true)
	stupid_self.AssertEqual(t, IsSequences(minSwapPlanA([]int{0, 7, 3, 10, 5}, []int{0, 2, 9, 4, 11})), true)
}

func TestIsSequences(t *testing.T) {
	stupid_self.AssertEqual(t, IsSequences([]int{1, 5, 3}), false)
	stupid_self.AssertEqual(t, IsSequences([]int{1}), true)
	stupid_self.AssertEqual(t, IsSequences([]int{1, 2}), true)
	stupid_self.AssertEqual(t, IsSequences([]int{4, 1}), false)
	stupid_self.AssertEqual(t, IsSequences([]int{1, 2, 3, 4, 5}), true)
}
