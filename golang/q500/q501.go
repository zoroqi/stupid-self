package q500

/**
# 501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-mode-in-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 愚笨的思考

`[1,null,2,2]`狗血的测试数据, 二叉树数组表示法还是错误的.

简单方案就是直接通过map进行统计. Simple方案, 两次提交耗时差4ms, 也是神奇了. orz

如何不适用map方案呢?

因为是BST, 二叉树的中序遍历是递增序列. 发现golang最小内存都有6044kb好大啊.
*/

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func FindMode(root *TreeNode) []int {
	r := make([]int, 0)
	max := 1
	pre := -1
	count := 1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if node.Val != pre {
			count = 1
			if count == max {
				r = append(r, node.Val)
			}
		} else {
			count++
			if count > max {
				r = append(r[:0], node.Val)
				max = count
			} else if count == max {
				r = append(r, node.Val)
			}
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return r
}

func FindMode_Simple(root *TreeNode) []int {
	m := make(map[int]int)
	FindMode_SimpleDfs(root, m)
	r := make([]int, 0)
	max := -1
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

func FindMode_SimpleDfs(node *TreeNode, m map[int]int) {
	if node == nil {
		return
	}
	m[node.Val]++
	FindMode_SimpleDfs(node.Right, m)
	FindMode_SimpleDfs(node.Left, m)
}
