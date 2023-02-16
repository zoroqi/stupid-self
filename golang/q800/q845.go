package q800

func longestMountainPlanB(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	l := len(arr)
	tops := make([]int, 0)
	for i := 1; i < l-1; i++ {
		if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
			tops = append(tops, i)
		}
	}
	ltop := len(tops)
	max := 0
	for i := 0; i < ltop; i++ {
		// left 存储了左山脚前一个元素
		left := tops[i] - 1
		for ; left >= 0; left-- {
			if arr[left] >= arr[left+1] {
				break
			}
		}
		// right 存储了右山脚
		right := tops[i] + 1
		for ; right < l-1; right++ {
			if arr[right] <= arr[right+1] {
				break
			}
		}
		// 这里 left 多算了一个距离, 这里就不用再 +1 了
		if max < right-left {
			max = right - left
		}
	}

	return max
}
