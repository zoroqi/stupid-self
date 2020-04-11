package stupid_self

import "fmt"

func PrintTwoDigitArray(a [][]int) {
	fmt.Println("-----------")
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("-----------")
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
