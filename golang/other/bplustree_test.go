package other

import (
	"fmt"
	"reflect"
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
			t.Fatalf("binsearch failed, k:%d, %d, wanted %d", k, binsearch(arr, k), v)
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
	for _, v := range arr {
		tree.Insert(v, v)
	}

	for _, v := range arr {
		if !tree.Insert(v, v+5) {
			t.Fatalf("insert failed, %d", v)
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
	//for i := 49; i >= 1; i -= 4 {
	//	tree.Insert(i, i)
	//}
	//for i := 50; i <= 200; i += 1 {
	//	tree.Insert(i, i)
	//}
	for i := 0; i < 15; i++ {
		tree.Insert(i, i)
	}
	treePrint(tree.(*bptree[int, int]))
	fmt.Println(tree.Search(28))
}

func TestSibling(t *testing.T) {
	bp := NewBPTree[int, int](5)
	tree := bp.(*bptree[int, int])
	for i := 0; i < 50; i++ {
		tree.Insert(i, i)
	}
	//treePrint(tree)
	node35 := tree.findNode(35)
	n35r := tree.rightSibling(node35)
	if !reflect.DeepEqual(n35r.keys, []int{37, 38}) {
		t.Fatalf("rightSibling failed, %v, wanted %v", n35r.keys, []int{37, 38})
	}
	n37l := tree.leftSibling(n35r)
	if !reflect.DeepEqual(n37l.keys, node35.keys) {
		t.Fatalf("leftSibling failed, %v, wanted %v", n37l.keys, node35.keys)
	}
	n37r := tree.rightSibling(n35r)
	if !reflect.DeepEqual(n37r.keys, []int{39, 40}) {
		t.Fatalf("rightSibling failed, %v, wanted %v", n37r.keys, []int{39, 40})
	}
	n39l := tree.leftSibling(n37r)
	if !reflect.DeepEqual(n39l.keys, n35r.keys) {
		t.Fatalf("leftSibling failed, %v, wanted %v", n39l.keys, n35r.keys)
	}
	node48 := tree.findNode(48)
	no := tree.rightSibling(node48)
	if no != nil {
		t.Fatalf("rightSibling failed, %v, wanted %v", no, nil)
	}
	node0 := tree.findNode(0)
	no = tree.leftSibling(node0)
	if no != nil {
		t.Fatalf("leftSibling failed, %v, wanted %v", no, nil)
	}
	p35r := tree.rightSibling(node35.parent)
	if !reflect.DeepEqual(p35r.keys, []int{39, 41}) {
		t.Fatalf("leftSibling failed, %v, wanted %v", p35r.keys, []int{39, 41})
	}
	p39l := tree.leftSibling(p35r)
	if !reflect.DeepEqual(p39l.keys, node35.parent.keys) {
		t.Fatalf("leftSibling failed, %v, wanted %v", p39l.keys, node35.parent.keys)
	}
}
