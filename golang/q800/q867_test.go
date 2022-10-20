package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestTranspose(t *testing.T) {
	stupid_self.AssertEqual(t, transpose(
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
		[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}})
	stupid_self.AssertEqual(t, transpose(
		[][]int{{1, 2, 3}, {4, 5, 6}}),
		[][]int{{1, 4}, {2, 5}, {3, 6}})
}
