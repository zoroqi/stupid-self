package q1900

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	stupid_self.AssertEqual(t, getConcatenationCopy([]int{}), []int{})
	stupid_self.AssertEqual(t, getConcatenation([]int{1, 2, 1}), []int{1, 2, 1, 1, 2, 1})
	stupid_self.AssertEqual(t, getConcatenation([]int{1, 3, 2, 1}), []int{1, 3, 2, 1, 1, 3, 2, 1})
}
