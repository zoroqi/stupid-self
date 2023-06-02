package q600

import (
	"fmt"
	. "github.com/zoroqi/stupid-self/golang"
)

func tree2str(root *TreeNode) string {
	var dfs func(n *TreeNode) string
	dfs = func(n *TreeNode) string {
		if n == nil {
			return ""
		}
		if n.Left != nil && n.Right != nil {
			return fmt.Sprintf("%d(%s)(%s)", n.Val, dfs(n.Left), dfs(n.Right))
		} else if n.Left != nil {
			return fmt.Sprintf("%d(%s)", n.Val, dfs(n.Left))
		} else if n.Right != nil {
			return fmt.Sprintf("%d()(%s)", n.Val, dfs(n.Right))
		} else {
			return fmt.Sprintf("%d", n.Val)
		}
	}
	return dfs(root)
}
