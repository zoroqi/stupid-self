package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestMostCommonWord(t *testing.T) {
	stupid_self.AssertEqual(t, mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.3", []string{"hit"}), "ball")
	stupid_self.AssertEqual(t, mostCommonWord("Bob hit a ball, the hit BA dLL  fledw far after it was hit. ", []string{"hit"}), "ball")
}
