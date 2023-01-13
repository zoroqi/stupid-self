package q800

import (
	"sort"
	"strings"
)

type dictnode struct {
	children map[rune]*dictnode
}

func (d *dictnode) add(s []rune) {
	l, tail := len(s), len(s)-1
	if l != 0 {
		if c, ok := d.children[s[tail]]; ok {
			c.add(s[:tail])
		} else {
			d.children[s[tail]] = &dictnode{children: make(map[rune]*dictnode)}
			d.children[s[tail]].add(s[:tail])
		}
	}
}
func (d *dictnode) leafDeep() []int {
	r := []int{}
	var dfs func(i int, n *dictnode)
	dfs = func(i int, n *dictnode) {
		if n == nil || len(n.children) == 0 {
			r = append(r, i)
		}
		for _, v := range n.children {
			dfs(i+1, v)
		}
	}
	dfs(0, d)
	return r
}

func minimumLengthEncoding(words []string) int {
	root := &dictnode{children: make(map[rune]*dictnode)}
	for _, v := range words {
		root.add([]rune(v))
	}
	leaf := root.leafDeep()
	s := 0
	for _, v := range leaf {
		s = s + v + 1
	}
	return s
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func minimumLengthEncodingPlanB(words []string) int {

	w := make([]string, len(words))
	for k, v := range words {
		w[k] = Reverse(v)
		//w[k] = v
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
		//sort.Slice(w, func(i, j int) bool {
		//	il, jl := len(w[i])-1, len(w[j])-1
		//	is, js := w[i], w[j]
		//	for ii, jj := il, jl; ii >= 0 && jj >= 0; ii, jj = ii-1, jj-1 {
		//		if is[ii] < js[jj] {
		//			return true
		//		} else if is[ii] > js[jj] {
		//			return false
		//		}
		//	}
		//	return il < jl
		//})
	})
	s := 0
	for i := 0; i < len(w)-1; i++ {
		if !strings.HasPrefix(w[i+1], w[i]) {
			s = s + len(w[i]) + 1
		}
	}
	s = s + len(w[len(w)-1]) + 1

	return s
}

func com(i int, j int, words []string) bool {
	il, jl := len(words[i])-1, len(words[j])-1
	is, js := []rune(words[i]), []rune(words[j])
	for ii, jj := il, jl; ii >= 0 && jj >= 0; ii, jj = ii-1, jj-1 {
		if is[ii] < js[jj] {
			return true
		} else if is[ii] > js[jj] {
			return false
		}
	}
	return il < jl
}
