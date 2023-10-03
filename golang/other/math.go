package other

func PascalsTriangle(line int) [][]int {
	if line <= 0 {
		return [][]int{}
	}

	r := make([][]int, line)
	for i := 0; i < line; i++ {
		r[i] = make([]int, i+1)
		r[i][0] = 1
		r[i][i] = 1
		for j := 1; j < i; j++ {
			r[i][j] = r[i-1][j-1] + r[i-1][j]
		}
	}
	return r
}

func GetPascalsTriangle(line int, index int) int {
	if line <= 0 {
		return 0
	}
	if index > line || index < 1 {
		return 0
	}
	if index == line || index == 1 {
		return 1
	}
	n := PascalsTriangle(line)
	return n[line-1][index-1]
	//n!/k!(n-k)!
	//factorial := func(i int) int {
	//	m := 1
	//	for j := 1; j <= i; j++ {
	//		m *= j
	//	}
	//	return m
	//}
	//l, i := line-1, index-1
	//return factorial(l) / (factorial(i) * factorial(l-i))
}
