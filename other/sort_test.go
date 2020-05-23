package other

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort([]int{9,8,3,4,5,6,7,8,9}))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{9,8,3,4,5,6,7,8,9}))
}

