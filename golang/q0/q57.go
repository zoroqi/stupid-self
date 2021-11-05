package q0

import (
	. "github.com/zoroqi/stupid-self"
)

func Insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) <= 0 {
		return [][]int{newInterval}
	}
	// 左右插入情况
	var r [][]int
	if intervals[0][0] > newInterval[1] {
		r = append(r, newInterval)
		return append(r, intervals...)
	}
	if intervals[len(intervals)-1][1] < newInterval[0] {
		return append(intervals, newInterval)
	}

	intersection := false
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			r = append(r, intervals[i])
		} else {
			if !intersection {
				intersection = true
				r = append(r, make([]int, 2))
				r[len(r)-1][0] = MinInts(intervals[i][0], intervals[i][1], newInterval[0], newInterval[1])
				// 直接插入右边界默认值
				r[len(r)-1][1] = newInterval[1]
			}
		}

		if intervals[i][0] > newInterval[1] {
			r = append(r, intervals[i])
		} else {
			if intersection {
				r[len(r)-1][1] = MaxInts(intervals[i][0], intervals[i][1], newInterval[0], newInterval[1])
			}
		}
	}
	return r
}
