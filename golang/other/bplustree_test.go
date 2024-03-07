package other

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func Test_binsearch(t *testing.T) {
	arr := []int{10, 20, 20, 30}
	d := map[int]int{
		5:  0,
		10: 0,
		15: 1,
		20: 1,
		25: 3,
		30: 3,
		35: 4,
	}
	for k, v := range d {
		if binsearch(arr, k) != v {
			t.Fatalf("binsearch failed, k:%d, %d, wanted %d", k, binsearch(arr, k), v)
		}
	}
}

func TestBptree_Delete2(t *testing.T) {
	for i := 0; i < 2000; i++ {
		TestbpTree(1000, nil, rand.Intn(20)+5, t)
	}
}

func TestBptree(t *testing.T) {
	arrs := [][]int{
		{7, 97, 90, 18, 86, 99, 91, 50, 11, 6, 28, 73, 88, 27, 95, 15, 79, 41, 98, 44, 29, 9, 5, 84, 69, 78, 39, 33, 16, 3, 71, 17, 56, 62, 40, 20, 2, 32, 1, 21, 63, 43, 81, 94, 80, 59, 64, 74, 47, 67, 87, 76, 96, 48, 61, 12, 53, 75, 89, 34, 42, 36, 30, 10, 77, 8, 93, 57, 25, 45, 51, 23, 37, 65, 22, 54, 49, 26, 66, 14, 46, 19, 58, 13, 31, 82, 85, 70, 24, 35, 52, 4, 83, 68, 92, 38, 72, 60, 55, 0},
		{64, 22, 51, 98, 73, 87, 99, 2, 14, 68, 70, 80, 79, 30, 65, 37, 86, 8, 6, 36, 41, 11, 15, 19, 3, 48, 97, 38, 44, 63, 88, 45, 55, 77, 84, 23, 78, 49, 31, 92, 90, 10, 5, 42, 58, 13, 1, 0, 74, 96, 66, 61, 12, 57, 24, 82, 76, 20, 33, 35, 40, 60, 50, 26, 46, 69, 93, 47, 83, 75, 27, 56, 67, 17, 34, 18, 43, 72, 91, 85, 7, 81, 32, 28, 54, 94, 25, 59, 29, 21, 52, 39, 95, 4, 9, 71, 16, 62, 53, 89},
		{39, 8, 47, 26, 32, 48, 49, 2, 28, 12, 0, 31, 6, 42, 20, 7, 3, 30, 21, 11, 44, 34, 5, 24, 29, 9, 4, 46, 36, 45, 10, 33, 22, 41, 17, 14, 19, 18, 35, 40, 37, 15, 13, 23, 27, 43, 25, 1, 16, 38},
		{20, 30, 12, 1, 37, 6, 23, 10, 11, 27, 38, 25, 24, 13, 18, 14, 5, 2, 0, 16, 39, 4, 15, 32, 28, 36, 9, 22, 17, 8, 35, 19, 21, 31, 26, 34, 7, 29, 33, 3},
		{12, 2, 24, 6, 26, 32, 4, 17, 34, 8, 5, 20, 29, 10, 30, 7, 16, 1, 13, 35, 3, 25, 28, 18, 33, 14, 9, 38, 36, 22, 31, 15, 39, 0, 19, 23, 37, 11, 21, 27},
		{16, 9, 26, 18, 28, 21, 20, 4, 19, 27, 6, 11, 5, 2, 25, 17, 10, 1, 12, 24, 22, 15, 0, 29, 8, 23, 3, 7, 14, 13},
		{20, 27, 5, 9, 10, 12, 0, 26, 2, 18, 8, 7, 21, 29, 16, 6, 11, 24, 1, 17, 14, 13, 28, 4, 22, 3, 19, 15, 25, 23},
		{16, 7, 27, 26, 13, 14, 0, 1, 21, 12, 8, 17, 25, 5, 2, 9, 19, 18, 22, 20, 29, 3, 24, 23, 15, 4, 6, 28, 11, 10},
		{15, 18, 12, 0, 28, 4, 21, 9, 1, 17, 27, 6, 26, 3, 8, 19, 7, 13, 14, 16, 23, 25, 20, 11, 22, 24, 29, 5, 10, 2},
		{39, 1, 15, 19, 3, 22, 2, 23, 0, 31, 36, 28, 17, 30, 12, 24, 34, 25, 13, 21, 11, 32, 18, 10, 35, 37, 4, 16, 8, 7, 5, 26, 29, 20, 14, 9, 6, 33, 27, 38},
	}
	for _, arr := range arrs {
		TestbpTree(len(arr), arr, 6, t)
	}

}

func TestbpTree(length int, arr []int, treeM int, t *testing.T) {
	if arr == nil {
		for i := 0; i < length; i++ {
			arr = append(arr, i)
		}
		rand.Shuffle(length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(treeM, arr)
			t.Fatal(err)
		}
	}()
	tree := NewBPTree[int, int](treeM)
	testCount := func(target int, key string) {
		count := 0
		tree.Iterator(func(k int, v int) error {
			count++
			return nil
		})

		if count != target {
			t.Fatalf("%s, count failed, wanted %d, %d", key, target, count)
		}
	}

	for _, v := range arr {
		if tree.Insert(v, v) {
			t.Fatalf("insert failed, %d", v)
		}
	}
	testCount(length, "insert")
	for _, v := range arr {
		if r, ok := tree.Search(v); !ok || r != v {
			t.Fatalf("search faild, ok: %t, wanted %d, %d", ok, v, r)
		}
	}
	for i, v := range arr[:length/2] {
		if !tree.Delete(v) {
			t.Fatalf("delete failed,%d %d", i, v)
		}
	}
	for _, v := range arr[length/2:] {
		if r, ok := tree.Search(v); !ok || r != v {
			t.Fatalf("search faild, ok: %t, wanted %d, %d", ok, v, r)
		}
	}
	for _, v := range arr[length/2:] {
		if tree.Delete(v + length) {
			t.Fatalf("delete failed, %d", v)
		}
	}
	testCount(length/2, "delete")
	for _, v := range arr {
		tree.Insert(v, v)
	}
	for _, v := range arr {
		if !tree.Insert(v, v+length) {
			t.Fatalf("insert failed, %d", v)
		}
	}
	for _, v := range arr {
		if r, ok := tree.Search(v); !ok || r != v+length {
			t.Fatalf("search faild, ok: %t, wanted %d, %d", ok, v+length, r)
		}
	}
	testCount(len(arr), "insert again")
	for i, v := range arr {
		if !tree.Delete(v) {
			t.Fatalf("delete failed,%d %d", i, v)
		}
	}
	testCount(0, "clean")
	for _, v := range arr {
		tree.Insert(v, v)
	}
	for i, v := range arr {
		if i%2 == 0 {
			if !tree.Insert(v, v+length) {
				t.Fatalf("insert failed, %d", v)
			}
		}
	}
	if err := tree.(*bptree[int, int]).walk(func(node *bpnode[int, int]) error {
		if node.leaf {
			if len(node.keys) > treeM {
				return fmt.Errorf("walk failed, leaf.keys > %d", treeM)
			}
		} else {
			if len(node.keys) > treeM-1 {
				return fmt.Errorf("walk failed, children.keys > %d", treeM-1)
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	tree := NewBPTree[int, int](10)
	for i := 0; i < 1000; i++ {
		tree.Insert(i, i)
	}
	treePrint(tree.(*bptree[int, int]))
	//fmt.Println(tree.Search(28))
}

func TestBptree_Delete(t *testing.T) {
	tree := NewBPTree[int, int](5)
	for i := 0; i < 20; i++ {
		tree.Insert(i, i)
	}
	tree.Delete(8)
}

func createBpTree(f, t int, tree BPTree[int, int]) BPTree[int, int] {
	for i := f; i < t; i++ {
		tree.Insert(i, i)
	}
	return tree
}

func Test_BptreeMerge(t *testing.T) {
	group := 0
	equals := func(key string, a, b any) {
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("%d: %s failed, %v, wanted %v", group, key, a, b)
		}
	}
	{
		group = 1
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node12 := tree.findNode(12)
		node14 := tree.findNode(14)
		newNode := tree.merge(node12, node14)
		equals("newLeaf", newNode.keys, []int{12, 13, 14, 15})
		equals("newLeafParent", newNode.parent.keys, []int{12, 16})
		equals("newLeafParentChldren1", newNode.parent.children[1].keys, []int{12, 13, 14, 15})
		equals("newLeafParentChldren2", newNode.parent.children[2].keys, []int{16, 17, 18, 19})
	}
	{
		group = 2
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node0 := tree.findNode(0)
		node4 := tree.findNode(4)
		newNode := tree.merge(node0.parent, node4.parent)
		equals("newNode", newNode.keys, []int{2, 4, 6})
		equals("newNode.children0", newNode.children[0].keys, []int{0, 1})
		equals("newNode.children1", newNode.children[1].keys, []int{2, 3})
		equals("newNode.children2", newNode.children[2].keys, []int{4, 5})
		equals("newNode.children3", newNode.children[3].keys, []int{6, 7})
		equals("newNode.parent", newNode.parent.keys, []int{2})
	}
	{
		group = 3
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node8 := tree.findNode(8)
		node12 := tree.findNode(12)
		// 当前的分裂逻辑导致经常会出现第一个节点是空的情况, 手动填入一个值
		node12.parent.children[0].keys = []int{11}
		node12.parent.children[0].values = []int{11}
		newNode := tree.merge(node8.parent, node12.parent)
		equals("newNode", newNode.keys, []int{8, 10, 11, 12, 14, 16})
		equals("newNode.children0", len(newNode.children[0].keys) == 0, true)
		equals("newNode.children1", newNode.children[1].keys, []int{8, 9})
		equals("newNode.children2", newNode.children[2].keys, []int{10, 11})
		equals("newNode.children3", newNode.children[3].keys, []int{11})
		equals("newNode.children4", newNode.children[4].keys, []int{12, 13})
		equals("newNode.children5", newNode.children[5].keys, []int{14, 15})
		equals("newNode.children6", newNode.children[6].keys, []int{16, 17, 18, 19})
		equals("newNode.parent", newNode.parent.keys, []int{8})
	}
	{
		group = 4
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node4 := tree.findNode(4)
		node6 := tree.findNode(6)
		node6.keys = []int{7, 8}
		newNode := tree.merge(node4, node6)
		equals("newLeaf", newNode.keys, []int{4, 5, 7, 8})
		equals("newLeafParent", newNode.parent.keys, []int{4})
		equals("newLeafParentChldren0", len(newNode.parent.children[0].keys) == 0, true)
		equals("newLeafParentChldren1", newNode.parent.children[1].keys, []int{4, 5, 7, 8})
	}
}

func Test_BpTreeBorrow(t *testing.T) {
	equals := func(key string, a, b any) {
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("%s failed, %v, wanted %v", key, a, b)
		}
	}
	{
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node0 := tree.findNode(0)
		l0, r0 := tree.nearSibling(node0)
		newNode := tree.borrow(l0, node0, r0)
		equals("newLeaf", newNode.keys, []int{0, 1, 2})
		equals("newLeafParent", newNode.parent.keys, []int{3})
		equals("newLeafParentChldren0", newNode.parent.children[0].keys, []int{0, 1, 2})
		equals("newLeafParentChldren1", newNode.parent.children[1].keys, []int{3})
	}

	{
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node2 := tree.findNode(2)
		l2, r2 := tree.nearSibling(node2)
		newNode := tree.borrow(l2, node2, r2)
		equals("newLeaf", newNode.keys, []int{1, 2, 3})
		equals("newLeafParent", newNode.parent.keys, []int{1})
		equals("newLeafParentChldren0", newNode.parent.children[0].keys, []int{0})
		equals("newLeafParentChldren1", newNode.parent.children[1].keys, []int{1, 2, 3})
	}

	{
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node2 := tree.findNode(2)
		l2, r2 := tree.nearSibling(node2.parent)
		newNode := tree.borrow(l2, node2.parent, r2)
		equals("newNode", newNode.keys, []int{2, 4})
		equals("newNodeParent", newNode.parent.keys, []int{6})
		equals("newNodeParentChldren0", newNode.parent.children[0].keys, []int{2, 4})
		equals("newNodeParentChldren1", newNode.parent.children[1].keys, []int{6})
	}
	{
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node12 := tree.findNode(12)
		l12, r12 := tree.nearSibling(node12.parent)
		newNode := tree.borrow(l12, node12.parent, r12)
		equals("newNode", newNode.keys, []int{10, 12, 14, 16})
		equals("newNodeParent", newNode.parent.keys, []int{8, 10})
		equals("newNodeChldren1", newNode.children[1].keys, []int{10, 11})
		equals("newNodeChldren2", newNode.children[2].keys, []int{12, 13})
		equals("newNodeChldren3", newNode.children[3].keys, []int{14, 15})
		equals("newNodeChldren4", newNode.children[4].keys, []int{16, 17, 18, 19})
	}
	{
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node2 := tree.findNode(2)
		l2, r2 := tree.nearSibling(node2.parent)
		r2.children[0].keys = []int{3}
		r2.children[0].values = []int{3}
		newNode := tree.borrow(l2, node2.parent, r2)
		equals("newNode", newNode.keys, []int{2, 3})
		equals("r2.keys", r2.keys, []int{4, 6})
		equals("parentKeys", newNode.parent.keys, []int{4})
	}
	{
		tree := createBpTree(0, 20, NewBPTree[int, int](5)).(*bptree[int, int])
		node4 := tree.findNode(4)
		l2, r2 := tree.nearSibling(node4.parent)
		node4.parent.children[0].keys = []int{3}
		node4.parent.children[0].values = []int{3}
		newNode := tree.borrow(l2, node4.parent, r2)
		equals("newNode", newNode.keys, []int{3, 4, 6})
	}
}
