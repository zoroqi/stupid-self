package q800

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func binaryGap(n int) int {
	before := -1
	i := 0
	m := n
	max := 0
	for m != 0 {
		if m%2 == 1 {
			if before != -1 {
				max = IntMax(i-before, max)
			}
			before = i
		}
		m = m >> 1
		i++
	}
	return max
}
