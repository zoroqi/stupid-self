package q400

import (
	"fmt"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	r := ReconstructQueue([][]int{{2, 4}, {3, 4}, {9, 0}, {0, 6}, {7, 1}, {6, 0}, {7, 3}, {2, 5}, {1, 1}, {8, 0}})
	fmt.Println(r)
}
