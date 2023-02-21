package other

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSubsequences(t *testing.T) {
	n := 6
	in := []int{}
	pow := 1
	dist := map[int]int{}
	cFunc := func(total, num int) int {
		a, b := 1, 1
		for i := 0; i < num; i++ {
			b *= i + 1
			a *= total - i
		}
		return a / b
	}
	for j := 0; j < n; j++ {
		in = append(in, j)
		pow = pow * 2
		dist[j] = cFunc(n, j)
	}
	dist[n] = 1
	sub := Subsequences(in)
	iter := IterSubsequences(in)
	itersub := [][]int{}
	for v, n := iter(); n; v, n = iter() {
		itersub = append(itersub, v)
	}
	if !reflect.DeepEqual(sub, itersub) {
		t.Fail()
	}
	// 数量
	if len(sub) != pow {
		t.Fatal("error total")
	}
	m := map[int]int{}
	dupMap := map[string]bool{}
	for _, v := range sub {
		dup := map[int]bool{}
		for _, v1 := range v {
			dup[v1] = true
			m[v1]++
		}
		dist[len(v)]--
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		key := ""
		for _, v1 := range v {
			key += strconv.Itoa(v1) + ","
		}
		dupMap[key] = true
		// 单个组合不能有重复的元素
		if len(dup) != len(v) {
			t.Fatal("error duplication")
		}
	}
	for _, v := range m {
		// 每个数字出现的次数是总结过的一半
		if v != pow/2 {
			t.Fatal("error size")
		}
	}
	// 去重数量
	if len(dupMap) != pow {
		t.Fatal("error total")
	}
	// 正确的分布
	for _, v := range dist {
		if v != 0 {
			t.Fatal("error distribution")
		}
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
		t.Fatal("error total")
	}
	m := map[string]bool{}
	for _, v := range f {
		// 单个排列需要保证数量正确
		if len(v) != n {
			t.Fatal("error len")
		}
		key := ""
		sumv := 0
		for _, v1 := range v {
			key += strconv.Itoa(v1) + ","
			sumv += v1
		}
		// 不能出现重复
		m[key] = true
		if sum != sumv {
			t.Fatal("error sum")
		}
	}
	// 排重后数量正确
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
