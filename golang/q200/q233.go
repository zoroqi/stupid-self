package q200

import (
	"math"
)

func countDigitOnePlanA(n int) int {
	count := func(i int) int {
		c := 0
		for i != 0 {
			if i%10 == 1 {
				c++
			}
			i /= 10
		}
		return c
	}
	c := 0
	for i := 0; i <= n; i++ {
		c += count(i)
	}
	return c
}

var countDigitOneCache = []int{}

const countDigitOneCacheLength = 9

func init() {
	countDigitOneCache = append(countDigitOneCache, 0)
	for i := 1; i <= countDigitOneCacheLength; i++ {
		countDigitOneCache = append(countDigitOneCache, i*int(math.Pow10(i-1)))
	}
}

func countDigitOnePlanB(n int) int {
	c09 := func(i int) int {
		if i <= 0 {
			return 0
		}
		if i <= countDigitOneCacheLength {
			return countDigitOneCache[i]
		}
		return i * int(math.Pow10(i-1))
	}

	length := func(i int) int {
		if i <= 0 {
			return 0
		}
		return int(math.Log10(float64(i))) + 1
	}

	var count func(n int, total int) int
	count = func(n int, total int) int {
		if n == 0 {
			return total
		}
		l := length(n) - 1
		base := int(math.Pow10(l))
		// 1000 ~ 1875
		if n/(base) == 1 {
			return count(n%base, total+c09(l)+n%base+1)
		}

		// 0000 ~ 3000
		// (n/base)*c09(l) + base
		return count(n%base, total+(n/base)*c09(l)+base)
	}

	return count(n, 0)
}
