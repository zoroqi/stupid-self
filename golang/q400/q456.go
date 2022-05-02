package q400

func find132patternPlanA(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[j] > (nums[i] + 1) {
				for k := j + 1; k < length; k++ {
					if nums[i] < nums[k] && nums[k] < nums[j] {
						return true
					}
				}
			}
		}
	}
	return false
}
