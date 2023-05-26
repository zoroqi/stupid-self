package q1900

func rotateGrid(grid [][]int, k int) [][]int {
	ml, nl := len(grid), len(grid[0])

	for startI := 0; startI < ml/2 && startI < nl/2; startI++ {
		m := ml - startI*2
		n := nl - startI*2
		totalLength := m*2 + n*2 - 4
		index := func(x, y int, r int) (int, int) {
			if r < m-1 {
				return x + r, y
			} else if r < m+n-2 {
				return x + m - 1, y + r - m + 1
			} else if r < m+n+m-3 {
				return x + m*2 + n - r - 3, y + n - 1
			} else {
				return x, y + n*2 + m*2 - r - 4
			}
		}
		road := make([]int, 0, totalLength)
		for j := 0; j < totalLength; j++ {
			ri, rj := index(startI, startI, j)
			road = append(road, grid[ri][rj])
		}

		rotate := k % totalLength
		road = append(road[totalLength-rotate:], road[:totalLength-rotate]...)
		for j := 0; j < totalLength; j++ {
			ri, rj := index(startI, startI, j)
			grid[ri][rj] = road[j]
		}
	}
	return grid

}
