package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestNumSpecialEquivGroups(t *testing.T) {
	data := [][]string{
		{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"},
		{"abc", "acb", "bac", "bca", "cab", "cba"},
		{"adadc", "bdbdb"},
	}
	for _, d := range data {
		stupid_self.AssertEqual(t, numSpecialEquivGroupsPlanA(d), numSpecialEquivGroupsPlanB(d))
	}
}
func TestNumSpecialEquivGroupsPlanA(t *testing.T) {
	stupid_self.AssertEqual(t, numSpecialEquivGroupsPlanA([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}), 3)
	stupid_self.AssertEqual(t, numSpecialEquivGroupsPlanA([]string{"abc", "acb", "bac", "bca", "cab", "cba"}), 3)
	stupid_self.AssertEqual(t, numSpecialEquivGroupsPlanA([]string{"adadc", "bdbdb"}), 2)
}
