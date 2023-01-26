package q500

import (
	. "github.com/zoroqi/stupid-self/golang"
	"math"
)

/**
# 508. 出现次数最多的子树元素和
```
给你一个二叉树的根结点，请你找出出现次数最多的子树元素和。一个结点的「子树元素和」定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。

你需要返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。



示例 1：
输入:

  5
 /  \
2   -3
返回 [2, -3, 4]，所有的值均只出现一次，以任意顺序返回所有值。

示例 2：
输入：

  5
 /  \
2   -5
返回 [2]，只有 2 出现两次，-5 只出现 1 次。



提示： 假设任意子树元素和均可以用 32 位有符号整数表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/most-frequent-subtree-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```
*/

// golang 简写声明+赋值操作真的会手抖引起作用域问题
func FindFrequentTreeSum(root *TreeNode) []int {

	m := make(map[int]int)
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := dfs(n.Left)
		r := dfs(n.Right)
		m[l+r+n.Val]++
		return l + r + n.Val
	}
	dfs(root)
	var r []int
	max := math.MinInt32
	for k, v := range m {
		if v > max {
			max = v
			r = r[:0]
			r = append(r, k)
		} else if v == max {
			r = append(r, k)
		}
	}
	return r
}
