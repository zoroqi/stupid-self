package q800

import (
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	ws1 := strings.Split(s1, " ")
	ws2 := strings.Split(s2, " ")
	counts := map[string]int{}
	for _, v := range ws1 {
		counts[v]++
	}
	for _, v := range ws2 {
		counts[v]++
	}
	var r []string
	for k, v := range counts {
		if v == 1 {
			r = append(r, k)
		}
	}
	return r
}
