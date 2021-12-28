package other

func MergeSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	c := l / 2
	a, b := nums[0:c], nums[c:]
	a = MergeSort(a)
	b = MergeSort(b)
	return merge(a, b)
}

func merge(a, b []int) []int {
	al, bl := len(a), len(b)
	ai, bi := 0, 0
	r := make([]int, 0, al+bl)
	for ai < al || bi < bl {
		for (bi >= bl && ai < al) || (ai < al && a[ai] <= b[bi]) {
			r = append(r, a[ai])
			ai++
		}

		for (ai >= al && bi < bl) || (bi < bl && b[bi] < a[ai]) {
			r = append(r, b[bi])
			bi++
		}
	}
	return r
}

// 好菜, 竟然些错了多个地方. 一处是笔误, 一处是没有想清楚.
func QuickSort(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}
func quickSort(nums []int, start, end int) []int {
	if end-start < 1 {
		return nums
	}
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
	quickSort(nums, start, s-1)
	quickSort(nums, s+1, end)
	return nums
}
