package q400

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

var q400list [][][]int

func init() {
	q400list = append(q400list, [][]int{{1, 2, 2, 2, 2, 7, 8, 9, 9, 10}, {3, 4, 5, 6}})
	q400list = append(q400list, [][]int{{1, 2, 3, 7, 8, 9, 9, 9, 9, 10}, {4, 5, 6}})
	q400list = append(q400list, [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {2, 3, 4, 5, 6, 7, 8, 9, 10}})
	q400list = append(q400list, [][]int{{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, {1, 2, 3, 4, 5, 6, 7, 8, 10}})
	q400list = append(q400list, [][]int{{2, 3, 3, 3, 3, 5, 5, 5, 7, 8}, {1, 4, 6, 9, 10}})
	q400list = append(q400list, [][]int{{1}, {}})
	q400list = append(q400list, [][]int{{1, 1}, {2}})
	q400list = append(q400list, [][]int{{2, 2}, {1}})
}

func TestFindDisappearedNumbersPlanB(t *testing.T) {
	for _, v := range q400list {
		stupid_self.AssertEqual(t, findDisappearedNumbersPlanB(v[0]), v[1])
	}
}

func TestFindDisappearedNumbersPlanC(t *testing.T) {
	for _, v := range q400list {
		stupid_self.AssertEqual(t, findDisappearedNumbersPlanC(v[0]), v[1])
	}
}

func TestBC(t *testing.T) {
	for _, v := range q400list {
		stupid_self.AssertEqual(t, findDisappearedNumbersPlanC(v[0]), findDisappearedNumbersPlanC(v[0]))
	}
}
