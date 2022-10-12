---
aliases:
- 814. 二叉树剪枝
- 814. binary tree pruning
tc:
- leetcode
- algorithm
leetcode:
  num: 814
  url: https://leetcode.cn/problems/binary-tree-pruning/
  tags:
  - 树
  - 二叉树
  - 深度优先遍历
id: 20221012213751_3b5aa704a1fa4206
date: "2022-10-12"
---

# 814. 二叉树剪枝

给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。

返回移除了所有不包含 1 的子树的原二叉树。

节点 node 的子树为 node 本身加上所有 node 的后代。

```
示例 1：

输入：root = [1,null,0, null, null,0,1]
输出：[1,null,0,null,null,null,1]
解释：
只有红色节点满足条件“所有不包含 1 的子树”。

示例 2:
输入：root = [1,0,1,0,0,0,1]
输出：[1,null,1,null,null,null,1]

示例 3:
输入：root = [1,1,0,1,1,0,1,0]
输出：[1,1,0,1,1,null,1]
```

**提示：**

* 树中节点的数目在范围 `[1, 200]` 内
* Node.val 为 0 或 1

## 愚笨的思考

这个题比较简单, 深度优先遍历就可以了, dfs 根据左右节点状态和自身 Val 判断自身状态.

最开始我返回的是一个 bool 在外层递归进行处理, 之后调整成返回 `*TreeNode` 这样判断少一些也容易进行处理.

### haskell 咋处理?

问题是如何删除数据呢? 主要是我要构建一个新的树, 而不是删除树的节点. 没有必要删除, 直接构建了一个新的树, 这样跟简单一些.

还是 `deriving` 好用啊, 简单就搞定了很多问题.
