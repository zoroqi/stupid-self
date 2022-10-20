package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestUncommonFromSentences(t *testing.T) {
	stupid_self.AssertEqualFunc(t, uncommonFromSentences("this apple is sweet", "this apple is sour"), []string{"sweet", "sour"}, stupid_self.SetEqual)
	stupid_self.AssertEqualFunc(t, uncommonFromSentences("apple apple", "banana"), []string{"banana"}, stupid_self.SetEqual)
	stupid_self.AssertEqualFunc(t, uncommonFromSentences("apple banana", "a banana"), []string{"a", "apple"}, stupid_self.SetEqual)
}
