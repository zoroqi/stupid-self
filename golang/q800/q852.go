package q800

func peakIndexInMountainArrayPlanA(arr []int) int {
	l := len(arr) - 1
	for i := 1; i < l; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}

func peakIndexInMountainArrayPlanB(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := (l + r) / 2
		if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func peakIndexInMountainArrayPlanC(arr []int) int {
	l, max, index := len(arr)-1, -1, -1
	for i := 1; i < l; i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	return index
}
