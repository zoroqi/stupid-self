package q100

import (
	. "github.com/zoroqi/stupid-self"
	"math"
)

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	imax, imin := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < 0 {
			imax, imin = imin, imax
		}
		imax = IntMax(imax*n, n)
		imin = IntMin(imin*n, n)

		max = IntMax(max, imax)
	}
	return max
}

func first(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := math.MinInt32
	product := 1
	minusNum := 1
	interval := 0

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		p := 0
		if n < 0 {
			product *= n
			p = product
			if minusNum > 0 {
				minusNum = product
			} else {
				interval = 0
				minusNum = 1
			}
		} else if n == 0 {
			if interval == 0 {
				p = 0
			}
			minusNum = 1
			product = 1
		} else {
			product *= n
			p = product
			interval++
		}
		max = IntMax(max, p)
	}

	if interval > 0 {
		max = IntMax(max, product/minusNum)
	}

	product = 1
	minusNum = 1
	interval = 0

	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		p := 0
		if n < 0 {
			product *= n
			p = product
			if minusNum > 0 {
				minusNum = product
			} else {
				interval = 0
				minusNum = 1
			}
		} else if n == 0 {
			if interval == 0 {
				p = 0
			}
			minusNum = 1
			product = 1
		} else {
			product *= n
			p = product
			interval++
		}
		max = IntMax(max, p)
	}
	if interval > 0 {
		max = IntMax(max, product/minusNum)
	}

	return max
}
