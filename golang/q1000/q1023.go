package q1000

import (
	"regexp"
)

func camelMatchPlanA(queries []string, pattern string) []bool {
	p := []byte("[a-z]*")
	s := make([]byte, 0, len(pattern)*7+2)
	s = append(s, '^')
	for i := 0; i < len(pattern); i++ {
		v := pattern[i]
		s = append(s, p...)
		s = append(s, v)

	}
	s = append(s, p...)
	s = append(s, '$')
	r := make([]bool, len(queries))
	reg := regexp.MustCompile(string(s))
	for i, v := range queries {
		r[i] = reg.MatchString(v)
	}
	return r
}
