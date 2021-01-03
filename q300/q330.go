package q300

import (
	"math"
)

func MinPatches(nums []int, n int) int {
	addNum := 0
	continuous := 0
	for _, v := range nums {
		if continuous >= n {
			break
		}
		for continuous < v-1 {
			if continuous >= n {
				break
			}
			continuous += continuous + 1
			addNum++
		}
		continuous += v
	}

	for continuous < n {
		continuous += continuous + 1
		addNum++
	}
	return addNum
}

// 第一版
func First_MinPatches(nums []int, n int) int {
	// 空集合可以直接计算
	if len(nums) <= 0 {
		return int(math.Log2(float64(n))) + 1
	}
	addNum := 0
	continuous := 0
	// 针对第一个数字不是1的情况, 计算需要添加多少个数字达nums[0]
	if nums[0] != 1 {
		addNum = int(math.Log2(float64(nums[0])))
		continuous = int(math.Pow(2, float64(addNum))) - 1
	}
	// 计算连续范围
	for _, v := range nums {
		if continuous >= n {
			break
		}
		for continuous < v-1 {
			if continuous >= n {
				break
			}
			continuous += continuous + 1
			addNum++
		}
		continuous += v
	}
	for continuous < n {
		continuous += continuous + 1
		addNum++
	}
	return addNum
}
