package q600

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func Test_findErrorNumsPlanA(t *testing.T) {
	stupid_self.AssertEqualFunc(t, findErrorNumsPlanB([]int{1, 2, 2, 4}), []int{2, 3}, stupid_self.SetEqual)
	stupid_self.AssertEqualFunc(t, findErrorNumsPlanB([]int{1, 1}), []int{1, 2}, stupid_self.SetEqual)
}
