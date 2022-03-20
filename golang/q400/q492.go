package q400

import "math"

func constructRectangle(area int) []int {
	c := int(math.Sqrt(float64(area)))
	for area%c != 0 {
		c--
	}
	return []int{area/c, c}
}
