package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestFindReplaceString(t *testing.T) {
	type td struct {
		s       string
		indices []int
		sources []string
		targets []string
		r       string
	}

	ds := []td{}
	ds = append(ds, td{"abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "fff"}, "eeebfff"})
	ds = append(ds, td{"abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "fff"}, "eeecd"})
	ds = append(ds, td{"abcd", []int{1, 2}, []string{"b", "ec"}, []string{"eee", "fff"}, "aeeecd"})
	ds = append(ds, td{"vmokgggqzp", []int{3, 5, 1}, []string{"kg", "ggq", "mo"}, []string{"s", "so", "bfr"}, "vbfrssozp"})
	ds = append(ds, td{"vmokgggqzp", []int{1, 3, 5}, []string{"mo", "kg", "ggq"}, []string{"bfr", "s", "so"}, "vbfrssozp"})
	ds = append(ds, td{"jjievdtjfb", []int{4, 6, 1}, []string{"md", "tjgb", "jf"}, []string{"foe", "oov", "e"}, "jjievdtjfb"})
	ds = append(ds, td{"jjievdtjfb", []int{1, 4, 6}, []string{"jf", "md", "tjgb"}, []string{"e", "foe", "oov"}, "jjievdtjfb"})
	ds = append(ds, td{"abcde", []int{2, 2}, []string{"cdef", "bc"}, []string{"f", "fe"}, "abcde"})

	tp := func(d td) (string, []int, []string, []string) {
		return d.s, d.indices, d.sources, d.targets
	}

	bc := func(d td) bool {
		b := findReplaceStringPlanB(tp(d))
		c := findReplaceStringPlanC(tp(d))
		return b == c && b == d.r
	}

	for _, d := range ds {
		stupid_self.AssertEqual(t, bc(d), true)
	}
}
