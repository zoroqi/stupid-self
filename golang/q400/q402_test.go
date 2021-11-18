package q400

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	//stupid_self.AssertEqual(t, RemoveKdigits("1432219", 3), "1219")
	//stupid_self.AssertEqual(t, RemoveKdigits("10200", 1), "200")
	//stupid_self.AssertEqual(t, RemoveKdigits("10", 2), "0")
	//stupid_self.AssertEqual(t, RemoveKdigits("10", 1), "0")
	stupid_self.AssertEqual(t, RemoveKdigits("112", 1), "11")
	//stupid_self.AssertEqual(t, RemoveKdigits("10200", 2), "0")
}
