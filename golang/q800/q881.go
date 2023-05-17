package q800

import (
	"sort"
)

func numRescueBoatsPlanA(people []int, limit int) int {
	if len(people) == 0 {
		return 0
	}

	sort.Ints(people)
	count := 0
	for i, j := 0, len(people)-1; i <= j; {
		if people[i]+people[j] <= limit {
			i++
			j--
		} else {
			j--
		}
		count++
	}
	return count
}
