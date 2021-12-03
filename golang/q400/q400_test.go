package q400

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestFindNthDigitPlanB(t *testing.T) {
	stupid_self.AssertEqual(t, FindNthDigitPlanB(11), 0)
	stupid_self.AssertEqual(t, FindNthDigitPlanB(1001), 7)
	stupid_self.AssertEqual(t, FindNthDigitPlanB(100143), 5)
}
