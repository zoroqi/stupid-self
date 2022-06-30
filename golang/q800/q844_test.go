package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestBackspaceComparePlanA(t *testing.T) {
	stupid_self.AssertEqual(t, backspaceComparePlanA("ab##", "c#d#"), true)
	stupid_self.AssertEqual(t, backspaceComparePlanA("ab##", ""), true)
	stupid_self.AssertEqual(t, backspaceComparePlanA("a#c", "b"), false)
}
