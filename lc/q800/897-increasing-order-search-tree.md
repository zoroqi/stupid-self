---
aliases:
- 897. 递增顺序搜索树
- 897. increasing order search tree
tc:
- leetcode
- algorithm
leetcode:
  num: 897
  url: https://leetcode.cn/problems/increasing-order-search-tree
  tags:
  - 二叉树
  - 栈
  - 深度优先遍历
date: "2023-05-13"
id: 20230513143738_741e6fbb89944ab5
---

给你一棵二叉搜索树的 root ，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。

```
示例 1：

输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]


示例 2：

输入：root = [5,1,7]
输出：[1,null,5,null,7]
```

提示：

* 树中节点数的取值范围是 [1, 100]
* 0 <= Node.val <= 1000


## 愚笨的思考

最开始想直接递归进行原地的替换, 但是实现上感觉有点费劲, 还是生成一个对应的中序遍历结果, 生成一个新的 tree 简单一点.

如果使用原地变换, 需要返回两个节点, 分别是最小节点和最大节点.
返回单个节点实现有点麻烦.

haskell 就不实现了, 决定把三种遍历实现一遍.

实现过程中, 一个很有意思的问题, 我是否要实现一个深度优先遍历
