package q600

import (
	"fmt"
	"sort"
)

func findLongestChainPlanA(pairs [][]int) int {
	if len(pairs) <= 1 {
		return len(pairs)
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	fmt.Println(pairs)
	l := len(pairs)
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				if len(pairs[j]) == 2 {
					pairs[j] = append(pairs[j], 1)
				}
				pairs[j][2]++
				pairs[j][1] = pairs[i][1]
			}
		}
	}
	fmt.Println(pairs)
	m := 1
	for i := 0; i < l; i++ {
		if len(pairs[i]) == 3 {
			if m < pairs[i][2] {
				m = pairs[i][2]
			}
		}
	}
	return m
}

func findLongestChainPlanB(pairs [][]int) int {
	if len(pairs) <= 1 {
		return len(pairs)
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	cache := map[int]map[int]int{}
	l := len(pairs)
	for i := 0; i < l; i++ {
		if _, ok := cache[pairs[i][1]]; !ok {
			cache[pairs[i][1]] = map[int]int{}
		}
		cache[pairs[i][1]][pairs[i][0]] = 1
	}
	c := 0
	for i := 0; i < l; i++ {
		for k, v := range cache {
			if k < pairs[i][0] {
				for kk, vv := range v {
					c++
					n := vv + 1
					if cache[pairs[i][1]][kk] < n {
						cache[pairs[i][1]][kk] = n
					}
				}
			}
		}
	}
	m := 0
	fmt.Println(c, "a")
	for _, v := range cache {
		for _, vv := range v {
			if m < vv {
				m = vv
			}
		}
	}

	return m
}

func findLongestChainPlanC(pairs [][]int) int {
	if len(pairs) <= 1 {
		return len(pairs)
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	c := 0
	l := len(pairs)
	cache := make(map[int]int)
	for i := 0; i < l; i++ {
		cache[pairs[i][1]] = 1
		for k, v := range cache {
			c++
			if k < pairs[i][0] {
				nv := v + 1
				if cache[pairs[i][1]] < nv {
					cache[pairs[i][1]] = nv
				}
			}
		}
	}
	m := 0
	for _, v := range cache {
		if m < v {
			m = v
		}
	}

	return m
}

func findLongestChainPlanD(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	n := len(pairs)
	dp := make([]int, n)
	c := 0
	for i, p := range pairs {
		dp[i] = 1
		for j, q := range pairs[:i] {
			c++
			if p[0] > q[1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}
