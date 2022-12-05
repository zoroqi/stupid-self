package q800

import (
	"math"
)

func consecutiveNumbersSum(n int) int {
	doubleN := 2 * n
	f := func(l int) bool {
		return (doubleN-l*l+l)%(2*l) == 0
	}
	lMax := int(math.Sqrt(float64(doubleN)))
	count := 0
	for l := 1; l <= lMax; l++ {
		if f(l) {
			count++
		}
	}
	return count
}
