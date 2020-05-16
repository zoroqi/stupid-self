package q0

import (
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestInsert(t *testing.T) {
	stupid_self.AssertEqual(t, Insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}),
		[][]int{{1, 2}, {3, 10}, {12, 16}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 2}, {3, 5}, {6, 7}, {9, 10}, {12, 16}}, []int{4, 8}),
		[][]int{{1, 2}, {3, 8}, {9, 10}, {12, 16}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 2}, {3, 4}}, []int{-3, -2}),
		[][]int{{-3, -2}, {1, 2}, {3, 4}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 2}, {3, 4}}, []int{5, 6}),
		[][]int{{1, 2}, {3, 4}, {5, 6}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 2}, {3, 4}}, []int{0, 6}),
		[][]int{{0, 6}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 2}, {7, 9}}, []int{4, 5}),
		[][]int{{1, 2}, {4, 5}, {7, 9}})

	stupid_self.AssertEqual(t, Insert([][]int{{1, 3}, {7, 9}}, []int{10, 11}),
		[][]int{{1, 3}, {7, 9}, {10, 11}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 3}, {7, 9}}, []int{-1, 0}),
		[][]int{{-1, 0}, {1, 3}, {7, 9}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 3}, {7, 9}}, []int{4, 6}),
		[][]int{{1, 3}, {4, 6}, {7, 9}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 3}, {7, 9}}, []int{2, 8}),
		[][]int{{1, 9}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 3}, {7, 9}}, []int{0, 10}),
		[][]int{{0, 10}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 3}, {7, 9}}, []int{4, 8}),
		[][]int{{1, 3}, {4, 9}})
	stupid_self.AssertEqual(t, Insert([][]int{{1, 3}, {7, 9}}, []int{2, 4}),
		[][]int{{1, 4}, {7, 9}})
}
