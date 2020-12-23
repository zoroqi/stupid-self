package other

import (
	"math/rand"
	"time"
)

// 猴子排序
// 验证 -> 交换两个元素 -> 验证 -> 交换 ....
func MonkeySort(arr []int) (a []int, loopCount int) {
	a = arr
	if isOrder(arr) {
		return
	}
	length := len(arr)
	rand.Seed(time.Now().UnixNano())
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	for !isOrder(arr) {
		loopCount++
		r1 := abs(rand.Int() % length)
		r2 := abs(rand.Int() % length)
		arr[r1], arr[r2] = arr[r2], arr[r1]
	}
	return
}

func isOrder(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
