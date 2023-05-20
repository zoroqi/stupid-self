package q800

import (
	"sort"
)

func numSpecialEquivGroupsPlanA(words []string) int {
	if len(words) <= 1 {
		return len(words)
	}

	nodiff := func(s string) string {
		l := len(s)
		odd := make([]byte, 0, (l+1)/2)
		even := make([]byte, 0, (l+1)/2)
		for i := 0; i < l; i++ {
			if i%2 == 0 {
				even = append(even, s[i])
			} else {
				odd = append(odd, s[i])
			}
		}
		sort.Slice(odd, func(i, j int) bool {
			return odd[i] < odd[j]
		})
		sort.Slice(even, func(i, j int) bool {
			return even[i] < even[j]
		})
		return string(odd) + "_" + string(even)
	}
	dup := map[string]int{}
	for _, v := range words {
		dup[nodiff(v)]++
	}

	return len(dup)
}

func numSpecialEquivGroupsPlanB(words []string) int {
	if len(words) <= 1 {
		return len(words)
	}

	nodiff := func(s string) string {
		l := len(s)
		odd := make([]byte, 26)
		even := make([]byte, 26)
		for i := 0; i < l; i++ {
			if i%2 == 0 {
				odd[s[i]-'a']++
			} else {
				even[s[i]-'a']++
			}
		}
		return string(odd) + "_" + string(even)
	}
	dup := map[string]int{}
	for _, v := range words {
		dup[nodiff(v)]++
	}

	return len(dup)
}
