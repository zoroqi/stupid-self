package other

func SimpleMedian(nums []int) int {
	n := QuickSort(nums)
	return n[len(n)/2]
}

// 基于快排实现. 中位数(n): n左侧所有数字均<=n, n右侧数字均>=n, n在整个数组的1/2处.
//快排的执行就是在寻找中间数字, 当标杆元素在1/2处就是中位数返回, 标杆在1/2左侧中位数在标杆右侧范围内,
//同理标杆在1/2右侧中位数在标杆左侧范围内. 通过这种方式可以减少对数组的全排序, 只需要做部分排序就可以找到中位数.
func QuickMedian(nums []int) int {
	index := len(nums) / 2
	return quickMedian(nums, index, 0, len(nums)-1)
}

func quickMedian(nums []int, index, start, end int) int {
	temp := nums[start]
	s, e := start, end
	for s < e {
		for s < e && nums[e] >= temp {
			e--
		}
		if s < e {
			nums[s] = nums[e]
		}
		for s < e && nums[s] <= temp {
			s++
		}
		if s < e {
			nums[e] = nums[s]
		}
	}
	nums[s] = temp
	if s == index {
		return nums[s]
	}
	if s > index {
		return quickMedian(nums, index, start, s-1)
	} else {
		return quickMedian(nums, index, s+1, end)
	}
}
