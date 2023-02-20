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

func TestPermutations(t *testing.T) {
	n := 3
	in := []int{}
	fac := 1
	sum := 0
	for j := 0; j < n; j++ {
		in = append(in, j)
		fac = fac * (j + 1)
		sum += j
	}
	f := Permutations(in)
	if len(f) != fac {
		t.Fail()
	}
	m := map[string]bool{}
	for _, v := range f {
		if len(v) != n {
			t.Fatal("error len")
		}
		key := ""
		sumv := 0
		for _, v1 := range v {
			key += strconv.Itoa(v1) + ","
			sumv += v1
		}
		m[key] = true
		if sum != sumv {
			t.Fatal("error sum")
		}
	}
	if len(m) != fac {
		t.Fatal("error total")
	}
}

func BenchmarkPermutations(b *testing.B) {
	n := 6
	for i := 3; i <= n; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			in := []int{}
			for j := 0; j < i; j++ {
				in = append(in, j)
			}
			for k := 0; k < b.N; k++ {
				Permutations(in)
			}
		})
	}
}
