package q100

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestReorderList(t *testing.T) {
	assert(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5})
	assert(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 8, 1, 7, 2, 6, 3, 5, 4})
	assert(t, []int{0, 1, 2}, []int{0, 2, 1})
	stupid_self.AssertEqual(t, []int{0, 1, 2}, []int{0, 1, 2})
}
func BenchmarkReorderList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := stupid_self.NewListNode([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		ReorderList(l)
	}
}

func assert(t *testing.T, arr []int, result []int) {
	l := stupid_self.NewListNode(arr)
	ReorderList(l)
	reorder := stupid_self.ListNodeToArray(l)
	for i := range result {
		if result[i] != reorder[i] {
			t.Errorf("err\n%v\n%v", result, reorder)
		}
	}
	t.Logf("success\n%v\n%v", arr, reorder)
}
