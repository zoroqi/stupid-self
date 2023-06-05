package q600

import "sort"

func findKthNumberPlanA(m int, n int, k int) int {
	if k <= 1 {
		return 1
	}
	if k >= m*n {
		return m * n
	}

	arr := make([]int, 0, m*n)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			arr = append(arr, i*j)
		}
	}
	sort.Ints(arr)
	return arr[k-1]
}

func findKthNumberPlanC(m int, n int, k int) int {
	c := map[int]int{}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			c[i*j]++
		}
	}
	arr := make([]int, 0, len(c))
	for k := range c {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	i := k
	for i > 0 {
		i -= c[arr[0]]
		if i <= 0 {
			return arr[0]
		}
		arr = arr[1:]
	}
	return arr[0]
}

func findKthNumberPlanD(m int, n int, k int) int {
	return sort.Search(m*n, func(x int) bool {
		count := x / n * n
		for i := x/n + 1; i <= m; i++ {
			count += x / i
		}
		return count >= k
	})
}
