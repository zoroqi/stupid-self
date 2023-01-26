package q400

import (
	"strings"
)

func RemoveKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	var nums []byte

	for i := range num {
		c := num[i]
		for k > 0 && len(nums) > 0 && c < nums[len(nums)-1] {
			nums = nums[:len(nums)-1]
			k--
		}
		nums = append(nums, c)
	}
	nums = nums[:len(nums)-k]
	n := strings.TrimLeft(string(nums), "0")
	if n == "" {
		return "0"
	}
	return n
}

func second402(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	nums := removeZero(rk([]rune(num), k))
	if len(nums) == 0 {
		return "0"
	}
	return string(nums)
}

func rk(nums []rune, k int) []rune {
	if k <= 0 || len(nums) == 0 || len(nums) <= k {
		return nums
	}
	index := findMax402(nums)
	nums = append(nums[0:index], nums[index+1:]...)
	return rk(nums, k-1)
}

func findMax402(nums []rune) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == '9' {
			return i
		}
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

func first402(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	nums := []rune(num)

	findMax := func(r []rune) int {
		max := r[0]
		index := 0
		for i := 0; i < len(r); i++ {
			if max < r[i] {
				index = i
				max = r[i]
			}
		}
		return index
	}
	tempK := k
	for i := 0; tempK > 0; i++ {
		removeIndex := findMax(nums[i:i+tempK]) + i
		nums = append(nums[0:removeIndex], nums[removeIndex+1:]...)
		tempK--
	}

	newNums := removeZero(nums)
	if len(newNums) == 0 {
		return "0"
	}
	return string(newNums)
}

func removeZero(nums []rune) []rune {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == '0' {
			zeroCount++
		} else {
			break
		}
	}
	return nums[zeroCount:]
}
