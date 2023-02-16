package other

import (
	"math"
)

func Subsequences[T any](arr []T) [][]T {
	l := len(arr)
	if l == 0 {
		return [][]T{}
	}
	initLen := int(math.Pow(2, float64(l)))
	if initLen < 0 {
		initLen = 100
	}
	r := make([][]T, 0, initLen)
	r = append(r, []T{})
	for i := 0; i < l; i++ {
		ll := len(r)
		for j := 0; j < ll; j++ {
			c := append([]T{}, r[j]...)
			r = append(r, append(c, arr[i]))
		}
	}
	return r
}

func IterSubsequences[T any](arr []T) func() ([]T, bool) {
	l := len(arr)
	initLen := int(math.Pow(2, float64(l)))
	if initLen < 0 {
		initLen = 100
	}

	r := make([][]T, 0, initLen)
	r = append(r, []T{})

	i := 0
	j := -1
	ll := len(r)
	end := false
	return func() ([]T, bool) {
		if end {
			return nil, false
		}
		if j == -1 {
			j++
			return []T{}, true
		}
		if j >= ll {
			j = 0
			ll = len(r)
			i++
		}
		if i >= l {
			end = true
			return nil, false
		}
		c := append([]T{}, r[j]...)
		d := append(c, arr[i])
		r = append(r, d)
		j++
		return d, i < l
	}
}
