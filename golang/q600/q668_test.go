package q600

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func Test_findKthNumberPlanA(t *testing.T) {
	stupid_self.AssertEqual(t, findKthNumberPlanA(3, 3, 5), 3)
	stupid_self.AssertEqual(t, findKthNumberPlanA(2, 3, 6), 6)
	stupid_self.AssertEqual(t, findKthNumberPlanA(1, 1, 1), 1)
	stupid_self.AssertEqual(t, findKthNumberPlanA(4, 4, 10), 6)
}

func Test_findKthNumberPlanD(t *testing.T) {
	type d struct {
		m, n, k int
	}
	data := []d{}
	data = append(data, d{3, 3, 5})
	data = append(data, d{2, 3, 6})
	data = append(data, d{1, 1, 1})
	data = append(data, d{4, 4, 10})
	for _, v := range data {
		stupid_self.AssertEqual(t, findKthNumberPlanD(v.m, v.n, v.k), findKthNumberPlanA(v.m, v.n, v.k))
	}
	stupid_self.AssertEqual(t, findKthNumberPlanD(9895, 28405, 100787757), 31666344)

}
