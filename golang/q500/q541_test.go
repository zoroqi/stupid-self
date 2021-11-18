package q500

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestReverseStr(t *testing.T) {
	stupid_self.AssertEqual(t, ReverseStr("abcd", 1), "abcd")
	stupid_self.AssertEqual(t, ReverseStr("abcd", 2), "bacd")
	stupid_self.AssertEqual(t, ReverseStr("abcdefgh", 3), "cbadefhg")
	stupid_self.AssertEqual(t, ReverseStr("abcdefg", 3), "cbadefg")
	stupid_self.AssertEqual(t, ReverseStr("abcdefghijk", 2), "bacdfeghjik")
}
