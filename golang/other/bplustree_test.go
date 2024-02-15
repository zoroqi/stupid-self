package other

import (
	"testing"
)

func Test_binsearch(t *testing.T) {
	arr := []int{10, 20, 30}
	d := map[int]int{
		5:  0,
		10: 0,
		15: 1,
		20: 1,
		25: 2,
		30: 2,
		35: 3,
	}
	for k, v := range d {
		if binsearch(arr, k) != v {
			t.Fatalf("binsearch failed, %d, %d", k, v)
		}
	}
}

func TestBptree(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	tree := NewBPTree[int, int](4)
	testCount := func(target int) {
		count := 0
		tree.Iterator(func(k int, v int) error {
			count++
			return nil
		})

		if count != target {
			t.Fatalf("count failed, wanted %d, %d", target, count)
		}
	}

	for _, v := range arr {
		if tree.Insert(v, v) {
			t.Fatalf("insert failed, %d", v)
		}
	}
	testCount(len(arr))
	for _, v := range arr {
		if r, ok := tree.Search(v); !ok || r != v {
			t.Fatalf("search faild, ok: %t, wanted %d, %d", ok, v, r)
		}
	}
	treePrint(tree.(*bptree[int, int]))
	//for _, v := range arr {
	//	if v < 50 {
	//		if !tree.Delete(v) {
	//			t.Fatalf("delete failed, %d", v)
	//		}
	//	} else {
	//		if tree.Delete(v + 5) {
	//			t.Fatalf("delete failed, %d", v)
	//		}
	//	}
	//}
	//testCount(5)
	//for _, v := range arr {
	//	tree.Insert(v, v)
	//}
	//
	for _, v := range arr {
		if !tree.Insert(v, v+5) {
			t.Fatalf("delete failed, %d", v)
		}
	}

	for _, v := range arr {
		if r, ok := tree.Search(v); !ok || r != v+5 {
			t.Fatalf("search faild, ok: %t, wanted %d, %d", ok, v+5, r)
		}
	}
	testCount(len(arr))
}

func TestInsert(t *testing.T) {
	tree := NewBPTree[int, int](5)
	for i := 0; i < 50; i += 4 {
		tree.Insert(i, i)
	}
	for i := 49; i >= 1; i -= 4 {
		tree.Insert(i, i)
	}
	for i := 50; i <= 200; i += 1 {
		tree.Insert(i, i)
	}
	treePrint(tree.(*bptree[int, int]))
}
