package q100

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	stupid_self.AssertEqual(t, MaxProduct([]int{1, 2, 0, 3, -1}), 3)
	stupid_self.AssertEqual(t, MaxProduct([]int{1, 2, 1, 3, -1}), 6)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 2, 1, 3, 1}), 6)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 0, 1, 3, 1}), 3)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 0, -1, 3, 1}), 3)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 0, -1, 0, 1}), 1)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1}), -1)
	stupid_self.AssertEqual(t, MaxProduct([]int{0}), 0)
	stupid_self.AssertEqual(t, MaxProduct([]int{0, 1}), 1)
	stupid_self.AssertEqual(t, MaxProduct([]int{1, 0}), 1)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 0}), 0)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 0, -1}), 0)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, -2}), 2)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 1}), 1)
	stupid_self.AssertEqual(t, MaxProduct([]int{-1, 1, 0}), 1)
	stupid_self.AssertEqual(t, MaxProduct([]int{-2, 0, -1}), 0)
	stupid_self.AssertEqual(t, MaxProduct([]int{2, -5, -2, -4, 3}), 24)
	stupid_self.AssertEqual(t, MaxProduct([]int{2, 2, 3, 0, 2, -1, 4, -2, 8, -3}), 192)
	stupid_self.AssertEqual(t, MaxProduct([]int{-2, -3, -3, -2, -3, -4}), 432)
	stupid_self.AssertEqual(t, MaxProduct([]int{1, -2, 1}), 1)
}
