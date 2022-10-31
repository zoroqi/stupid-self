package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestScoreOfParenthesesPlanA(t *testing.T) {
	stupid_self.AssertEqual(t, scoreOfParenthesesPlanA("()"), 1)
	stupid_self.AssertEqual(t, scoreOfParenthesesPlanA("(())"), 2)
	stupid_self.AssertEqual(t, scoreOfParenthesesPlanA("()()"), 2)
	stupid_self.AssertEqual(t, scoreOfParenthesesPlanA("(()(()))"), 6)
	stupid_self.AssertEqual(t, scoreOfParenthesesPlanA("(()(()()))"), 10)
	stupid_self.AssertEqual(t, scoreOfParenthesesPlanA("((()())(()))"), 12)
}
