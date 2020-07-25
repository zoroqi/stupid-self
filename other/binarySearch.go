package other

func BinarySearch(arr []int, num int) (index int) {
	index = -1
	if len(arr) == 0 {
		return
	}
	l, r := 0, len(arr)-1
	for r >= l {
		i := (r + l) / 2
		if arr[i] == num {
			index = i
			break
		} else if arr[i] > num {
			r = i
		} else {
			// 这里需要除法是向下取整,需要执行`i+1`
			l = i + 1
		}
	}
	return
}
