package q800

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	var monotonic func(n1, n2 int) bool

	l := len(nums) - 1
	for i := 0; i < l; i++ {
		if nums[i] != nums[i+1] {
			if monotonic != nil {
				if !monotonic(nums[i], nums[i+1]) {
					return false
				}
			} else {
				if nums[i] > nums[i+1] {
					monotonic = func(n1, n2 int) bool {
						return n1 > n2
					}
				} else {
					monotonic = func(n1, n2 int) bool {
						return n1 < n2
					}
				}
			}
		}
	}
	return true
}
