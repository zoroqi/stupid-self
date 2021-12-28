package q400

import (
	"sort"
)

func findDisappearedNumbersPlanB(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	sort.Ints(nums)
	length := len(nums)
	length2 := length + 1
	i, j := 1, 0
	var result []int
	for j < length && i < length2 {
		if nums[j] == i {
			i++
			j++
		} else {
			for j < length && i < length2 && nums[j] < i {
				k := nums[j]
				j++
				for ; j < length && i < length2 && nums[j] == k; j++ {
					if k == i {
						i++
					}
				}
				for k = k + 1; k < i; k++ {
					result = append(result, k)
				}
			}
			for j < length && i < length2 && nums[j] > i {
				for ; nums[j] > i && i < length2; i++ {
					result = append(result, i)
				}
			}
		}
	}
	if i < length2 {
		k := nums[j-1] + 1
		for ; k < length2; k++ {
			result = append(result, k)
		}
	}
	return result
}

func findDisappearedNumbersPlanC(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	var result []int
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		for n > 0 {
			n, nums[n-1] = nums[n-1], -n
		}
		if nums[i] > 0 && nums[nums[i]-1] <= 0 {
			nums[i] = 0
		}
	}
	for i, v := range nums {
		if v >= 0 {
			result = append(result, i+1)
		}
	}
	return result
}
