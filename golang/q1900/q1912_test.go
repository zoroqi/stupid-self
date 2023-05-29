package q1900

import (
	"fmt"
	"testing"
)

func TestMovieRentingSystem(t *testing.T) {
	// ["MovieRentingSystem","search","rent","rent","report","drop","search"]
	// [[3,[[0,1,5],[0,2,6],[0,3,7],[1,1,4],[1,2,7],[2,1,5]]],[1],[0,1],[1,2],[],[1,2],[2]]
	system := Q1912Constructor(3, [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}})
	fmt.Println(system.Search(1))
	system.Rent(0, 1)
	system.Rent(1, 2)
	fmt.Println(system.Report())
	system.Drop(1, 2)
	fmt.Println(system.Search(2))
}
