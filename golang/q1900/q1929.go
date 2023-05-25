package q1900

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func getConcatenationCopy(nums []int) []int {
	l := len(nums)
	r := make([]int, l*2)
	copy(r, nums)
	copy(r[l:], nums)
	return r
}
