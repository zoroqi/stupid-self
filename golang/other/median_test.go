package other

import (
	"fmt"
	"testing"
)

func TestSimpleMedian(t *testing.T) {
	v := []int{10, 1, 2, 3, 9, 4, 1, 2, 33, 12, 12, 12, 12, 2, 1, 1, 2, 2, 1, 2, 6, 45, 1, 4, 56, 23, 4, 8, 7, 34}
	//fmt.Println(v)
	fmt.Println(QuickMedian(v))
	//fmt.Println(v)
	fmt.Println(SimpleMedian(v))
	fmt.Println(v)
}
