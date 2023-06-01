package q600

func findErrorNumsPlanA(nums []int) []int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	l := len(nums)
	lose := -1
	dup := -1
	for i := 1; i <= l; i++ {
		switch m[i] {
		case 1:
		case 0:
			lose = i
		case 2:
			dup = i
		}
	}
	return []int{dup, lose}
}

func findErrorNumsPlanB(nums []int) []int {
	l := len(nums)
	m := make([]int, l+1)
	for _, v := range nums {
		m[v]++
	}
	lose := -1
	dup := -1
	for i := 1; i <= l; i++ {
		switch m[i] {
		case 1:
		case 0:
			lose = i
		case 2:
			dup = i
		}
	}
	return []int{dup, lose}
}
