package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestConsecutiveNumbersSum(t *testing.T) {
	stupid_self.AssertEqual(t, consecutiveNumbersSum(5), 2)
	stupid_self.AssertEqual(t, consecutiveNumbersSum(9), 3)
	stupid_self.AssertEqual(t, consecutiveNumbersSum(15), 4)
	stupid_self.AssertEqual(t, consecutiveNumbersSum(45), 6)
}
