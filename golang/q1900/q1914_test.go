package q1900

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestRotateGrid(t *testing.T) {
	stupid_self.AssertEqual(t,
		rotateGrid([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 2),
		[][]int{{3, 4, 8, 12}, {2, 11, 10, 16}, {1, 7, 6, 15}, {5, 9, 13, 14}})
}
