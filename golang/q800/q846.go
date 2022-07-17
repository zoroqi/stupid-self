package q800

import (
	"sort"
)

func isNStraightHandPlanA(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	m := map[int]int{}
	for _, v := range hand {
		m[v]++
	}
	arr := make([]int, 0, len(m))
	for k, _ := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for _, v := range arr {
		n := m[v]
		if n == 0 {
			continue
		}
		for i := v; i < v+groupSize; i++ {
			m[i] -= n
			if m[i] < 0 {
				return false
			}
		}
	}
	return true
}
