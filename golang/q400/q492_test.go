package q400

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestConstructRectanglePlanb(t *testing.T) {
	stupid_self.AssertEqual(t, constructRectangle(37), []int{37, 1})
	stupid_self.AssertEqual(t, constructRectangle(2), []int{2, 1})
	stupid_self.AssertEqual(t, constructRectangle(110), []int{11, 10})
	stupid_self.AssertEqual(t, constructRectangle(48), []int{8, 6})
}
