package q400

import (
	"fmt"
	"sort"
)

func ReconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return people
	}
	pl := len(people)

	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		ii := people[i]
		jj := people[j]
		if ii[0] == jj[0] {
			return ii[1] < jj[1]
		}
		return ii[0] < jj[0]
	})
	result := first(people, pl)
	return result
}

func first(people [][]int, pl int) [][]int {
	result := make([][]int, pl)
	var n int
	var c int
	var i int
	init := func(hi int) {
		n = hi
		c = -1
		i = 0
	}
	init(people[0][0])
	for _, v := range people {
		if n != v[0] {
			init(v[0])
		}
		j := v[1]
		for i < pl {
			if result[i] == nil {
				c++
			}
			if c == j {
				result[i] = v
				i++
				break
			}
			i++
		}
	}
	return result
}
