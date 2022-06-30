package q800

func backspaceComparePlanA(s string, t string) bool {
	m := func(str string) string {
		r := make([]rune, 0, len(str))
		for _, v := range str {
			if v == '#' {
				if len(r) > 0 {
					r = r[:len(r)-1]
				}
			} else {
				r = append(r, v)
			}
		}
		return string(r)
	}
	return m(s) == m(t)
}
