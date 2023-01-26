package other

func Compilers_11_42a(n int, z, w []int, x, y [][]int) {
	for i := 0; i < n; i++ {
		z[i] = z[i] / w[i]
		for j := 0; j < n; j++ {
			x[i][j] = y[i][j] * y[i][j]
			z[j] = z[j] + x[i][j]
		}
	}
}

func Compilers_11_42c(n int, z, w []int, x, y [][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x[i][j] = y[i][j] * y[i][j]
		}
	}
	for i := 0; i < n; i++ {
		z[i] = z[i] / w[i]
		for j := 0; j < n; j++ {
			z[j] = z[j] + x[i][j]
		}
	}
}
func Compilers_11_42d(n int, z, w []int, x, y [][]int) {
	for i := 0; i < n; i++ {
		z[i] = z[i] / w[i]
		for j := 0; j < n; j++ {
			x[i][j] = y[i][j] * y[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			z[j] = z[j] + x[i][j]
		}
	}
}
