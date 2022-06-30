package q800

func canVisitAllRoomsPlanA(rooms [][]int) bool {
	in := make([]bool, len(rooms))
	total := 0
	var queue []int
	add := func(i int) {
		if in[i] {
			return
		}
		in[i] = true
		total++
		queue = append(queue, i)
	}
	empty := func() bool {
		return len(queue) == 0
	}
	remove := func() int {
		l := queue[0]
		queue = queue[1:]
		return l
	}
	add(0)
	for !empty() {
		for _, v := range rooms[remove()] {
			add(v)
		}
	}
	return total >= len(rooms)
}

func canVisitAllRoomsPlanB(rooms [][]int) bool {
	in := make([]bool, len(rooms))
	in[0] = true
	var dfs func(int, [][]int, []bool, int) int
	dfs = func(number int, rooms [][]int, in []bool, total int) int {
		for _, v := range rooms[number] {
			if !in[v] {
				in[v] = true
				total = dfs(v, rooms, in, total) + 1
			}
		}
		return total
	}
	return dfs(0, rooms, in, 1) >= len(rooms)
}
