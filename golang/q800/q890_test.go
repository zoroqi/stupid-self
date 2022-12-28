package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestFindAndReplacePattern(t *testing.T) {
	stupid_self.AssertEqualFunc(t, findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"), []string{"mee", "aqq"}, stupid_self.SetEqual)
	stupid_self.AssertEqualFunc(t, findAndReplacePattern([]string{"abcd", "deqe", "meem", "aqqf", "dkda", "ccce"}, "abba"), []string{"meem"}, stupid_self.SetEqual)
}
