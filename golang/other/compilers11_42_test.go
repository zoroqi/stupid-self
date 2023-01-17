package other

import (
	"testing"
)

func buildParam(n int) ([]int, []int, [][]int, [][]int) {
	z := initArr(n, inc)
	w := initArr(n, inc)
	x := initArr(n, func(i int) []int {
		return initArr(n, inc)
	})
	y := initArr(n, func(i int) []int {
		return initArr(n, inc)
	})
	return z, w, x, y
}

func inc(i int) int {
	return i + 1
}

func initArr[T any](length int, value func(int) T) []T {
	a := make([]T, length)
	for i := 0; i < length; i++ {
		a[i] = value(i)
	}
	return a
}

const arr_length = 10

func BenchmarkCompilers_11_42a(b *testing.B) {
	z, w, x, y := buildParam(arr_length)
	for i := 0; i < b.N; i++ {
		Compilers_11_42a(arr_length, z, w, x, y)
	}
}

func BenchmarkCompilers_11_42c(b *testing.B) {
	z, w, x, y := buildParam(arr_length)
	for i := 0; i < b.N; i++ {
		Compilers_11_42c(arr_length, z, w, x, y)
	}
}
