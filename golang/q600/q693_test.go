package q600

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestHasAlternatingBits(t *testing.T) {
	stupid_self.AssertEqual(t, HasAlternatingBits(10), true)
	stupid_self.AssertEqual(t, HasAlternatingBits(6), false)
	stupid_self.AssertEqual(t, HasAlternatingBits(5), true)
	stupid_self.AssertEqual(t, HasAlternatingBits(1), true)
	stupid_self.AssertEqual(t, HasAlternatingBits(2), true)
	stupid_self.AssertEqual(t, HasAlternatingBits(11), false)
	stupid_self.AssertEqual(t, HasAlternatingBits(8), false)
}
