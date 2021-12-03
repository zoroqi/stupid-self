package q400

import (
	"math"
)

func FindNthDigitPlanB(n int) int {
	c := 1
	r := 0
	for m := n; ; {
		rr := c * 9 * pow(10, c-1)
		m = m - rr
		if m <= 0 {
			break
		}
		c++
		r += rr
	}
	m := n - r - 1
	num := m/c + pow(10, c-1)
	index := c - m%c - 1
	return indexNum(num, index)
}

func indexNum(f, i int) int {
	nums := make([]int, 0)
	for f > 0 {
		nums = append(nums, f%10)
		f = f / 10
	}
	return nums[i]
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
