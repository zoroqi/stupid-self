package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestLargeGroupPositions(t *testing.T) {
	stupid_self.AssertEqual(t, largeGroupPositions("abbxxxxzzy"), [][]int{{3, 6}})
	stupid_self.AssertEqual(t, largeGroupPositions("abc"), [][]int{})
	stupid_self.AssertEqual(t, largeGroupPositions("abcdddeeeeaabbbcd"), [][]int{{3, 5}, {6, 9}, {12, 14}})
	stupid_self.AssertEqual(t, largeGroupPositions("aba"), [][]int{})
	stupid_self.AssertEqual(t, largeGroupPositions("aaaaaa"), [][]int{{0, 5}})
	stupid_self.AssertEqual(t, largeGroupPositions("aaaabbbb"), [][]int{{0, 3}, {4, 7}})
	stupid_self.AssertEqual(t, largeGroupPositions("a"), [][]int{})
	stupid_self.AssertEqual(t, largeGroupPositions(""), [][]int{})
}
