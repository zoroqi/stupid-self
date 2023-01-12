package stupid_self

import (
	"fmt"
	"reflect"
	"testing"
)

func PrintTwoDigitArray(a [][]int) {
	fmt.Println("-----------")
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("-----------")
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func MaxInts(arr ...int) int {
	max := arr[0]
	for _, a := range arr {
		max = MaxInt(a, max)
	}
	return max

}

func MinInts(arr ...int) int {
	min := arr[0]
	for _, a := range arr {
		min = MinInt(a, min)
	}
	return min
}

func AssertEqual(t *testing.T, a, b interface{}) {

	var equals func(a, b interface{}) bool
	aType := reflect.TypeOf(a)
	switch aType.Kind() {
	case reflect.Struct:
		equals = DefaultEqual
	case reflect.Array, reflect.Slice:
		equals = ArrEqual
	default:
		equals = DefaultEqual
	}

	AssertEqualFunc(t, a, b, equals)
}

func AssertEqualFunc(t *testing.T, a, b interface{}, equals func(a, b interface{}) bool) {
	if a == nil || b == nil {
		t.Errorf("error: a or b is nil,\na:%v\nb:%v", a, b)
	}
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		t.Errorf("error: the type of a is not b,\na:%v\nb:%v", a, b)
	}

	if equals(a, b) {
		t.Logf("equal:\na:%v\nb:%v", a, b)
	} else {
		t.Errorf("noEqual:\na:%v\nb:%v", a, b)
	}
}

func DefaultEqual(a, b interface{}) bool {
	aType := reflect.TypeOf(a)
	switch aType.Kind() {
	case reflect.Array, reflect.Slice:
		return ArrEqual(a, b)
	default:
		return a == b
	}
}

func ArrEqual(a, b interface{}) bool {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	al := av.Len()
	bl := bv.Len()

	if al != bl {
		return false
	}
	for i := 0; i < al; i++ {
		if !DefaultEqual(av.Index(i).Interface(), bv.Index(i).Interface()) {
			return false
		}
	}
	return true
}

func SetEqual(a, b any) bool {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	al := av.Len()
	bl := bv.Len()
	if al != bl {
		return false
	}
	am := map[any]int{}
	for i := 0; i < al; i++ {
		iv := av.Index(i).Interface()
		am[iv]++
	}
	bm := map[any]int{}
	for i := 0; i < bl; i++ {
		iv := bv.Index(i).Interface()
		bm[iv]++
	}
	if len(am) != len(bm) {
		return false
	}
	for k := range am {
		if bm[k] == 0 {
			return false
		}
	}
	return true
}
func TreeNodeEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	av, ok := a.(*TreeNode)
	if !ok {
		return false
	}
	bv, ok := b.(*TreeNode)
	if !ok {
		return false
	}
	if av == nil && bv == nil {
		return true
	}
	if av == nil || bv == nil {
		return false
	}
	if av.Val == bv.Val {
		return TreeNodeEqual(av.Left, bv.Left) && TreeNodeEqual(av.Right, bv.Right)
	} else {
		return false
	}
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func PrintTreeNodeT(root *TreeNode, t *testing.T) {
	if root == nil {
		return
	}
	t.Logf("parent:%pIndex ,%+v\n", root, *root)
	PrintTreeNodeT(root.Left, t)
	PrintTreeNodeT(root.Right, t)
}
