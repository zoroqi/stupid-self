package q500

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	stupid_self.AssertEqual(t, ConvertToBase7(100), "202")
	stupid_self.AssertEqual(t, ConvertToBase7(-8), "-11")
}
