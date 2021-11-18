package q100

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	stupid_self.AssertEqual(t, TitleToNumber("A"), 1)
	stupid_self.AssertEqual(t, TitleToNumber("B"), 2)
	stupid_self.AssertEqual(t, TitleToNumber("AB"), 28)
	stupid_self.AssertEqual(t, TitleToNumber("ZY"), 701)
}
