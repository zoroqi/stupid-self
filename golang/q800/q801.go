package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
)

func minSwap(nums1 []int, nums2 []int) int {
	return minSwapPlanB(nums1, nums2)
}

func minSwapPlanA(nums1 []int, nums2 []int) ([]int, []int) {
	i := 0
	l := len(nums1) - 1
	for i < l {
		x1, x2 := nums1[i], nums1[i+1]
		y1, y2 := nums2[i], nums2[i+1]

		if x1 >= x2 || y1 >= y2 {
			nums1[i+1], nums2[i+1] = y2, x2
		}
		i++
	}
	return nums1, nums2
}

func IsSequences(numss ...[]int) bool {
	for _, nums := range numss {
		for i := 0; i < (len(nums) - 1); i++ {
			if nums[i] >= nums[i+1] {
				return false
			}
		}
	}
	return true
}

func minSwapPlanB(nums1 []int, nums2 []int) int {
	l := len(nums1)
	dp0, dp1 := 0, 1
	for i := 1; i < l; i++ {
		x1, x2 := nums1[i-1], nums1[i]
		y1, y2 := nums2[i-1], nums2[i]
		if x1 < x2 && y1 < y2 && x1 < y2 && y1 < x2 {
			dp0 = stupid_self.IntMin(dp0, dp1)
			dp1 = stupid_self.IntMin(dp0, dp1) + 1
		} else if x1 < y2 && y1 < x2 {
			dp0, dp1 = dp1, dp0+1
		} else {
			dp0, dp1 = dp0, dp1+1
		}
	}
	return stupid_self.IntMin(dp0, dp1)
}
