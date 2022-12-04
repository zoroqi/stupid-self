package q800

func numBusesToDestinationPlanA(routes [][]int, source int, target int) int {

	if source == target {
		return 0
	}
	index := map[int][]struct {
		line     int
		stations []int
	}{}

	for k, v := range routes {
		for _, vv := range v {
			index[vv] = append(index[vv], struct {
				line     int
				stations []int
			}{line: k, stations: routes[k]})
		}
	}

	queue := make([]struct {
		i        int
		transfer int
	}, 0, len(index))

	lineDup := map[int]bool{}
	appendQ := func(i int, transfer int) bool {
		for _, v := range index[i] {
			if lineDup[v.line] {
				continue
			}
			lineDup[v.line] = true
			for _, k := range v.stations {
				if k == target {
					return true
				}
				queue = append(queue, struct {
					i        int
					transfer int
				}{i: k, transfer: transfer})
			}
		}
		return false
	}
	if appendQ(source, 1) {
		return 1
	}
	for len(queue) != 0 {
		a := queue[0]
		if appendQ(a.i, a.transfer+1) {
			return a.transfer + 1
		}
		queue = queue[1:]
	}
	return -1
}
