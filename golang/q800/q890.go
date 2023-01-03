package q800

func findAndReplacePattern(words []string, pattern string) []string {
	p := []rune(pattern)
	if len(p) == 1 {
		return words
	}
	r := []string{}

	caesar, clean := func() (func(s, t rune) bool, func()) {
		mapper := [26]rune{}
		exist := [26]bool{}
		return func(s, t rune) bool {
				s -= 97
				t -= 97
				if mapper[s] == -1 && !exist[t] {
					mapper[s] = t
					exist[t] = true
					return true
				} else {
					return mapper[s] == t
				}
			},
			func() {
				for i := range mapper {
					mapper[i] = -1
					exist[i] = false
				}
			}
	}()

Outer:
	for _, w := range words {
		clean()
		for i, b := range w {
			if !caesar(b, p[i]) {
				continue Outer
			}

		}
		r = append(r, w)
	}
	return r
}
