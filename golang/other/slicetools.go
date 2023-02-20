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
			c := append(make([]T, 0, len(r[j])+1), r[j]...)
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

func Permutations[T any](arr []T) [][]T {
	l := len(arr)

	if l == 0 {
		return [][]T{}
	}
	initLen := 1
	for i := 1; i <= l; i++ {
		initLen *= i
	}

	r := make([][]T, 0, initLen)
	r = append(r, make([]T, 0, l))
	for i := 0; i < l; i++ {
		ll := len(r)
		for j := 0; j < ll; j++ {
			rl := len(r[j]) + 1
			kr := r[j]
			for k := 0; k < rl; k++ {
				a := make([]T, 0, l)
				a = append(a, kr[:k]...)
				a = append(a, arr[i])
				a = append(a, kr[k:]...)
				if k == 0 {
					r[j] = a
				} else {
					r = append(r, a)
				}
			}
		}
	}
	return r
}
