package q500

import (
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestFindRotateSteps1(t *testing.T) {
	stupid_self.AssertEqual(t, FindRotateSteps1("abcdef", "a"), 1)
	stupid_self.AssertEqual(t, FindRotateSteps1("abcdef", "c"), 3)
	stupid_self.AssertEqual(t, FindRotateSteps1("abcdefg", "d"), 4)
	stupid_self.AssertEqual(t, FindRotateSteps1("ab", "b"), 2)
}

func TestFindRotateSteps2(t *testing.T) {
	stupid_self.AssertEqual(t, FindRotateSteps1("abcdef", "ab"), 3)
	stupid_self.AssertEqual(t, FindRotateSteps1("abcdef", "cb"), 5)
	stupid_self.AssertEqual(t, FindRotateSteps1("abcdefg", "dg"), 8)
	stupid_self.AssertEqual(t, FindRotateSteps1("nyngl", "yyynnnnnnlllggg"), 19)
	stupid_self.AssertEqual(t, FindRotateSteps1("nyngl", "yyy"), 4)
	stupid_self.AssertEqual(t, FindRotateSteps1("nynlg", "yyyng"), 8)
	stupid_self.AssertEqual(t, FindRotateSteps1("abcdefghigklmnopqrstuvwxyzabcdefghigklmnopqrstuvwxyzabcdefghigklmnopqrstuvwxyzabcdefghigklmnopqrstuvwxyz", "hapqrdiencjaleidngjeiajslkefijcjwleicjafadf"), 100)
}
