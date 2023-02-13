package q800

import (
	"sort"
	"strings"
)

func findReplaceStringPlanA(s string, indices []int, sources []string, targets []string) string {
	if len(indices) == 0 {
		return s
	}
	r := strings.Builder{}
	indicesL := len(indices)
	l := len(s)
	if indices[0] != 0 {
		r.WriteString(s[0:indices[0]])
	}

	for index, v := range indices {
		source := sources[index]
		sl := len(source)
		match := false
		if v+sl <= l {
			match = s[v:v+sl] == source
		}
		if match {
			r.WriteString(targets[index])
			if index != indicesL-1 {
				r.WriteString(s[v+sl : indices[index+1]])
			} else {
				r.WriteString(s[v+sl:])
			}
		} else {
			if index != indicesL-1 {
				r.WriteString(s[v:indices[index+1]])
			} else {
				r.WriteString(s[indices[index]:])
			}
		}
	}
	return r.String()
}

func findReplaceStringPlanB(s string, indices []int, sources []string, targets []string) string {
	if len(indices) == 0 {
		return s
	}
	type st struct {
		i int
		s string
		t string
	}
	sts := make([]st, len(indices))
	for i, v := range indices {
		sts[i] = st{i: v, s: sources[i], t: targets[i]}
	}
	sort.Slice(sts, func(i, j int) bool {
		return sts[i].i < sts[j].i
	})
	r := strings.Builder{}
	indicesL := len(indices)
	if sts[0].i != 0 {
		r.WriteString(s[0:sts[0].i])
	}
	l := len(s)
	for index, v := range sts {
		source := v.s
		sl := len(source)
		match := false
		if v.i+sl <= l {
			match = s[v.i:v.i+sl] == source
		}
		if match {
			r.WriteString(v.t)
			if index != indicesL-1 {
				r.WriteString(s[v.i+sl : sts[index+1].i])
			} else {
				r.WriteString(s[v.i+sl:])
			}
		} else {
			if index != indicesL-1 {
				r.WriteString(s[v.i:sts[index+1].i])
			} else {
				r.WriteString(s[sts[index].i:])
			}
		}
	}

	return r.String()
}

func findReplaceStringPlanC(s string, indices []int, sources []string, targets []string) string {
	if len(indices) == 0 {
		return s
	}

	type ind struct {
		i int
		s string
		t string
	}
	l := len(s)
	inds := make([]ind, len(indices))
	for i, v := range indices {
		inds[i] = ind{i: v, s: sources[i], t: targets[i]}
	}
	sort.Slice(inds, func(i, j int) bool {
		return inds[i].i > inds[j].i
	})

	for _, v := range inds {
		if v.i+len(v.s) > l {
			continue
		}
		if s[v.i:v.i+len(v.s)] == v.s {
			s = s[:v.i] + v.t + s[v.i+len(v.s):]
		}
	}
	return s
}
