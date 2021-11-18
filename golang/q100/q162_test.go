package q100

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	// 官方实例
	stupid_self.AssertEqual(t, FindPeakElement([]int{1, 2, 1, 3, 5, 6, 4}), 5)
	stupid_self.AssertEqual(t, FindPeakElement([]int{1, 2, 3, 1}), 2)
	// 单调升
	stupid_self.AssertEqual(t, FindPeakElement([]int{1, 2}), 1)
	stupid_self.AssertEqual(t, FindPeakElement([]int{1, 2, 3}), 2)
	// 单调降
	stupid_self.AssertEqual(t, FindPeakElement([]int{2, 1}), 0)
	stupid_self.AssertEqual(t, FindPeakElement([]int{3, 2, 1}), 0)
	// 中间拐点, 中间最高
	stupid_self.AssertEqual(t, FindPeakElement([]int{2, 3, 1}), 1)
	// 中间拐点, 中间最低. 两边都是峰值
	stupid_self.AssertEqual(t, FindPeakElement([]int{3, 1, 2}), 2)
}
