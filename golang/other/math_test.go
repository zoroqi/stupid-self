package other

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestPascalsTriangle(t *testing.T) {
	stupid_self.AssertEqual(t, PascalsTriangle(0), [][]int{})
	stupid_self.AssertEqual(t, PascalsTriangle(1), [][]int{{1}})
	stupid_self.AssertEqual(t, PascalsTriangle(2), [][]int{{1}, {1, 1}})
	stupid_self.AssertEqual(t, PascalsTriangle(4), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}})
}

func TestGetPascalsTriangle(t *testing.T) {
	stupid_self.AssertEqual(t, GetPascalsTriangle(4, 1), 1)
	stupid_self.AssertEqual(t, GetPascalsTriangle(4, 2), 3)
	stupid_self.AssertEqual(t, GetPascalsTriangle(4, 3), 3)
	stupid_self.AssertEqual(t, GetPascalsTriangle(4, 4), 1)
	stupid_self.AssertEqual(t, GetPascalsTriangle(5, 4), 4)
}

func BenchmarkGetPascalsTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPascalsTriangle(50, 35)
		GetPascalsTriangle(70, 45)
	}
}
