package q0

import (
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestIsNumber(t *testing.T) {
	stupid_self.AssertEqual(t, IsNumber(".."), false)
	stupid_self.AssertEqual(t, IsNumber("0"), true)
	stupid_self.AssertEqual(t, IsNumber(" 0"), true)
	stupid_self.AssertEqual(t, IsNumber(" 0.1 "), true)
	stupid_self.AssertEqual(t, IsNumber("abc"), false)
	stupid_self.AssertEqual(t, IsNumber("1 a"), false)
	stupid_self.AssertEqual(t, IsNumber("2e10"), true)
	stupid_self.AssertEqual(t, IsNumber(" -90e3   "), true)
	stupid_self.AssertEqual(t, IsNumber(" 1e"), false)
	stupid_self.AssertEqual(t, IsNumber("e3"), false)
	stupid_self.AssertEqual(t, IsNumber(" 6e-1"), true)
	stupid_self.AssertEqual(t, IsNumber(" 99e2.5 "), false)
	stupid_self.AssertEqual(t, IsNumber("53.5e93"), true)
	stupid_self.AssertEqual(t, IsNumber(" --6 "), false)
	stupid_self.AssertEqual(t, IsNumber("-+3"), false)
	stupid_self.AssertEqual(t, IsNumber("95a54e53"), false)
	stupid_self.AssertEqual(t, IsNumber(".1"), true)
	stupid_self.AssertEqual(t, IsNumber("-.1"), true)
	stupid_self.AssertEqual(t, IsNumber("1."), true)
	stupid_self.AssertEqual(t, IsNumber(" 005047e+6"), true)
	stupid_self.AssertEqual(t, IsNumber("46.e3"), true)
	stupid_self.AssertEqual(t, IsNumber(".2e81"), true)
	stupid_self.AssertEqual(t, IsNumber("."), false)
	stupid_self.AssertEqual(t, IsNumber("+e"), false)
	stupid_self.AssertEqual(t, IsNumber(".e1"), false)
	stupid_self.AssertEqual(t, IsNumber("-."), false)
	stupid_self.AssertEqual(t, IsNumber("+."), false)
	stupid_self.AssertEqual(t, IsNumber("-e58"), false)
	stupid_self.AssertEqual(t, IsNumber("3.e3"), true)
}
