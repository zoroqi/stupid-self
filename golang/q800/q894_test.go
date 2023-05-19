package q800

import (
	"sort"
	"testing"
)

func TestAllPossibleFBT(t *testing.T) {
	for i := 1; i <= 20; i++ {
		b := allPossibleFBTPlanB(i)
		c := allPossibleFBTPlanC(i)
		if len(b) != len(c) {
			t.Errorf("len(b) != len(c) %d %d", len(b), len(c))
		}
		bs := []string{}
		for _, v := range b {
			bs = append(bs, v.String())
		}
		sort.Strings(bs)
		cs := []string{}
		for _, v := range c {
			cs = append(cs, v.String())
		}
		sort.Strings(cs)
		for i, v := range bs {
			if v != cs[i] {
				t.Errorf("v != cs[i] %s %s", v, cs[i])
			}
		}
	}
}
