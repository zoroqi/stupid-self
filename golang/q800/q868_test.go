package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestBinaryGap(t *testing.T) {
	stupid_self.AssertEqual(t, binaryGap(22), 2)
	stupid_self.AssertEqual(t, binaryGap(8), 0)
	stupid_self.AssertEqual(t, binaryGap(5), 2)
	stupid_self.AssertEqual(t, binaryGap(0), 0)
}
