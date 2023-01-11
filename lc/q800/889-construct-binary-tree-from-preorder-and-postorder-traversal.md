---
aliases:
- 889. 根据前序和后序遍历构造二叉树
- 889. construct binary tree from preorder and postorder traversal
tc:
- leetcode
- algorithm
leetcode:
  num: 889
  url: https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
  tags:
  - 二叉树
date: "2023-01-11"
id: 20230111054557_b7848f653712417a
---

# 889. 根据前序和后序遍历构造二叉树

给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵树的后序遍历，重构并返回二叉树。

如果存在多个答案，您可以返回其中 任何 一个。

```
示例：

输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]

示例：
输入: preorder = [1], postorder = [1]
输出: [1]
```

提示：**

* `1 <= preorder.length <= 30`
* `1 <= preorder[i] <= preorder.length`
* `preorder` 中所有值都 **不同**
* `postorder.length == preorder.length`
* `1 <= postorder[i] <= postorder.length`
* `postorder` 中所有值都 **不同**
* 保证 `preorder` 和 `postorder` 是同一棵二叉树的前序遍历和后序遍历

## 愚笨的思考

读题的时候的一个重要疑问是"为什么会有多个答案", 我上学的时候好像是"只要知道任何两个遍历顺序就可以知道二叉树的结构".
仔细想了想发现是不行的, 必须包括"中序"才可以实现.
因为当一个节点只有一个子节点的时候无法知道, 是无法知道这个子节点是左节点还是右节点.

最开始想的方式是, 先生成一个左偏树(前序遍历), 在根据后续遍历的内容进行调整.
但是在思考具体的执行过程中, 发现有太多的判断无法确定, 所以就放弃了.

只能换一个新的思路, 在思考前序和后序生成的顺序的时候,
    发现前序的第二个元素一定是左子树树根, 后序的倒数第二个元素是右子树根, 我可以通过这种方式递归进行生成.
针对是无法确定左右的情况的时候, 默认左子节点.

思路上是没有什么大问题的, 但实现起来一堆小问题.
主要是在计算偏移量的时候各种小错误, 好烦.

但是这个算法用 haskell 就不是很好实现了, 主要是需要处理偏移量, 看看能不能再想想别的办法吧.
或者学习使用 Data.Array 库看看, 这个库可能会有更好的操作效果.
