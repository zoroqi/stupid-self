package other

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSubsequences(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	sub := Subsequences(in)
	iter := IterSubsequences(in)
	itersub := [][]int{}
	for v, n := iter(); n; v, n = iter() {
		itersub = append(itersub, v)
	}
	if !reflect.DeepEqual(sub, itersub) {
		t.Fail()
	}
}

func BenchmarkSubsequences(b *testing.B) {
	n := 10
	for i := 3; i <= n; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			in := []int{}
			for j := 0; j < i; j++ {
				in = append(in, j)
			}
			for k := 0; k < b.N; k++ {
				Subsequences(in)
			}
		})
	}
}
