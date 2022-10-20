package q800

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	words := splitRegex.Split(strings.ToLower(paragraph), -1)
	bans := map[string]bool{}
	bans[""] = true
	for _, b := range banned {
		bans[b] = true
	}

	counts := map[string]int{}
	for _, w := range words {
		counts[w]++
	}
	max := -1
	r := ""
	for s, c := range counts {
		if c > max {
			if !bans[s] {
				r = s
				max = c
			}
		}
	}
	return r
}

var splitRegex = regexp.MustCompile("[ !?',;.]+")
