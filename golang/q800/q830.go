package q800

func largeGroupPositions(s string) [][]int {
	r := [][]int{}
	l := 0
	begin := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if l == 0 {
				begin = i - 1
			}
			l++
		} else {
			if l >= 2 {
				r = append(r, []int{begin, i - 1})
			}
			l = 0
		}
	}
	if l >= 2 {
		r = append(r, []int{begin, len(s) - 1})
	}
	return r
}
