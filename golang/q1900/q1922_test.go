package q1900

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestCountGoodNumbers(t *testing.T) {
	stupid_self.AssertEqual(t, countGoodNumbers(1), 5)
	stupid_self.AssertEqual(t, countGoodNumbers(4), 400)
	stupid_self.AssertEqual(t, countGoodNumbers(50), 564908303)
}
