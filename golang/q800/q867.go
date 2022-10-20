package q800

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	w := len(matrix)
	h := len(matrix[0])
	newMatrix := make([][]int, h)
	for i := 0; i < h; i++ {
		newMatrix[i] = make([]int, w)
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			newMatrix[j][i] = matrix[i][j]
		}
	}
	return newMatrix
}
