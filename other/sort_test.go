package other

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort([]int{9, 8, 3, 4, 5, 6, 7, 8, 9}))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{10, 1, 2, 3, 9, 4, 1, 2, 33, 323, 11, 45,  56, 23, 4, 8, 7, 34}))
}
