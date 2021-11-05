package q100

func FindPeakElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if (len(nums)) == 1 {
		return 0
	}
	index := 0
	l, r := 0, len(nums)-1
	for r-l > 0 {
		b := (l + r) / 2
		if nums[b] > nums[b+1] {
			index = b
			r = b
		} else {
			index = b + 1
			l = b + 1
		}
	}
	return index
}
