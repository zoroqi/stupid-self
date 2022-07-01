package q800

import "strings"

func toGoatLatinPlanA(sentence string) string {
	ss := strings.Split(sentence, " ")
	vote := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	ma := "ma"
	a := func() string {
		ma += "a"
		return ma
	}

	for i, s := range ss {
		r := []rune(s)
		if vote[r[0]] {
			s += a()
		} else {
			s = s[1:] + s[:1]
			s += a()
		}
		ss[i] = s
	}
	return strings.Join(ss, " ")
}
