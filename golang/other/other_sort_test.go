package other

import (
	"fmt"
	"testing"
)

func TestMonkeySort(t *testing.T) {
	fmt.Println(MonkeySort([]int{1, 2, 5, 2, 0, 1, 4}))
}

func TestStoogeSort(t *testing.T) {
	fmt.Println(StoogeSort([]int{1, 2, 5, 0, 2, 1, 4}))
}

func TestBeadSort(t *testing.T) {
	fmt.Println(BeadSort([]uint{9, 1, 2, 5, 0, 2, 1, 4, 20, 10}))
}
