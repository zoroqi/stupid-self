package q400

import "math"

func ArrangeCoins(n int) int {
	// total = (1 + n) * n / 2
	// total * 2 = n + n^2
	// n^2 + n - total * 2 = 0
	// n = (-1 + sqrt(1 + 8 * total))/2
	x := (-1.0 + math.Sqrt(1.0+8.0*float64(n))) / 2
	return int(x)
}
