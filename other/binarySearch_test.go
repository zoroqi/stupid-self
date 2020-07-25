package other

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	//sort := QuickSort([]int{10, 1, 2, 3, 9, 4, 1, 2, 33, 323, 11, 45,  56, 23, 4, 8, 7, 34})
	sort := QuickSort([]int{1,2,3,4})
	fmt.Println(sort)
	fmt.Println(BinarySearch(sort,1))
	fmt.Println(BinarySearch(sort,2))
	fmt.Println(BinarySearch(sort,3))
	fmt.Println(BinarySearch(sort,4))
	fmt.Println(BinarySearch(sort,5))
}
