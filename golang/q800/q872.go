package q800

import (
	. "github.com/zoroqi/stupid-self/golang"
	"reflect"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	data := func() (func(i int), func() []int) {
		a := []int{}
		return func(i int) {
				a = append(a, i)
			}, func() []int {
				return a
			}
	}
	var dfs func(n *TreeNode, app func(i int))
	dfs = func(n *TreeNode, app func(i int)) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			app(n.Val)
			return
		}
		dfs(n.Left, app)
		dfs(n.Right, app)
	}
	a1, r1 := data()
	dfs(root1, a1)
	a2, r2 := data()
	dfs(root2, a2)
	return reflect.DeepEqual(r1(), r2())
}
