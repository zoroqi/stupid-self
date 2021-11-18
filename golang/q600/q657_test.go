package q600

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestJudgeCircle(t *testing.T) {
	stupid_self.AssertEqual(t, JudgeCircle("DDD"), false)
	stupid_self.AssertEqual(t, JudgeCircle("UD"), true)
	stupid_self.AssertEqual(t, JudgeCircle("ULDR"), true)
}
