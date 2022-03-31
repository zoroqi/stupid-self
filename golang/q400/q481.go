package q400

func levelProgression(level int) []int {
	nums := []int{1, 2}
	for i := 0; i < level; i++ {
		nums = nextProgression(nums)
	}
	return nums
}

func nextProgression(nums []int) []int {
	var n []int
	n = append(n, 1)
	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			if n[len(n)-1] == 1 {
				n = append(n, 2)
			} else {
				n = append(n, 1)
			}
		} else {
			if n[len(n)-1] == 1 {
				n = append(n, 2)
				n = append(n, 2)
			} else {
				n = append(n, 1)
				n = append(n, 1)
			}
		}
	}
	return n
}

func magicalString(length int) int {
	nums := infiniteProgression(length)
	count := 0
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			count++
		}
	}
	return count
}

func infiniteProgression(length int) []int {
	if length <= 0 {
		return nil
	}
	if length <= 3 {
		return [][]int{{1}, {1, 2}, {1, 2, 2}}[length-1]
	}
	nums := make([]int, 3, length+2)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 2
	//before := false
	for i := 2; len(nums) < length; i++ {
		beforeNum := nums[len(nums)-1]
		//beforeNum := 1
		//if before {
		//	before = false
		//} else {
		//	beforeNum = 2
		//	before = true
		//}
		queueNum := nums[i]
		if queueNum == 1 {
			if beforeNum == 1 {
				nums = append(nums, 2)
			} else {
				nums = append(nums, 1)
			}
		} else {
			if beforeNum == 1 {
				nums = append(nums, 2)
				nums = append(nums, 2)
			} else {
				nums = append(nums, 1)
				nums = append(nums, 1)
			}
		}
	}
	return nums[0:length]
}
