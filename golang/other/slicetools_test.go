package other

import (
	"reflect"
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
