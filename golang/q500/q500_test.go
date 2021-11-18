package q500

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestFindWords(t *testing.T) {
	stupid_self.AssertEqual(t,FindWords([]string{"Hello", "Alaska", "Dad", "Peace"}),[]string{"Alaska", "Dad"})
}