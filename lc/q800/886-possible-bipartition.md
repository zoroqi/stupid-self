---
aliases:
- 886. 可能的二分法
- 886. possible bipartition
tc:
- leetcode
- algorithm
leetcode:
  num: 886
  url: https://leetcode.cn/problems/possible-bipartition/
  tags:
  - 深度优先遍历
  - 并查集
  - 图
date: "2023-02-27"
id: 20230227203510_14cd5ed49d274d1a
---

# 886. 可能的二分法

给定一组 `n` 人（编号为 `1, 2, ..., n`）， 我们想把每个人分进**任意**大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。

给定整数 `n` 和数组 `dislikes` ，其中 `dislikes[i] = [ai, bi]` ，表示不允许将编号为 `ai` 和  `bi`的人归入同一组。当可以用这种方法将所有人分进两组时，返回 `true`；否则返回 `false`。

```
示例 1：

输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
输出：true
解释：group1 [1,4], group2 [2,3]
示例 2：

输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
输出：false
示例 3：

输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
输出：false
```

**提示：**

* `1 <= n <= 2000`
* `0 <= dislikes.length <= 10^4`
* `dislikes[i].length == 2`
* `1 <= dislikes[i][j] <= n`
* `ai < bi`
* `dislikes` 中每一组都 **不同**

## 愚笨的思考

### 单纯的直觉 PlanA

第一反应是, 把所有节点划分成至少两个独立的图就可以了.

* 创建两个组 g1 和 g2
* 从 1 遍历到 n
    * 把当前节点加入一个组中, 不喜欢的节点加入另一个组中
    * 判断另一个组是否可以添加
        * 不可以直接返回 false
* 返回 true

这个直觉的方案有一个问题, 在出现一个可以同时选择两个分组的情况, 可能会出现回溯的情况, 这个比较麻烦.

### 尝试递归 PlanB

在 PlanA 的基础上, 调整成递归进行分组. 加入到分组中, 并没有别的处理.

使用深度优先进行递归加入.

最后的耗时 400+ , 看看怎么优化看看.

仔细想了想, 那个选择分组的问题不存在, 我把 dislike 构建的映射从有向变成了无向.
当出现可以随记分组的, 那就是一个独立的图, 加入哪一个分组都是可以的.

#### 优化的 PlanC

线一条道走到黑, 继续这个方向上走.

发现之前有很多无用的判断, 递归中是可以干掉的, 调整存储结构, 耗时直接减少到了 130+.

用最快的代码测试, 多次结果在 90ms 左右, 那我还有 30% 的优化空间.

但是感觉不能在这个方向上继续了, 需要用别的实现了.

对最快的代码进行进行分析, 对空间的使用, 要更好, 对分组的使用也更巧妙.
只用一个 visited 就实现了分组判断, 我用的是两个 map 进行的. 果然没有 map 速度会好很多.

### 并查集的方案 PlanD

这是标签提供的一种实现, 但完全不知道这个是什么, 所以看看如何进行实现.
代码就是讲解中的代码, 简单改造了一下.

* [并查集 - OI Wiki](https://oi-wiki.org/ds/dsu/)
* [并查集 - wiki](https://zh.wikipedia.org/wiki/%E5%B9%B6%E6%9F%A5%E9%9B%86)

## 并查集

并查集是一种数据结构，用于处理一些**不交集**（Disjoint sets，一系列**没有重复元素**的集合）的**合并及查询问题**.

通过构建一棵树实现, 所有节点包含一个指向父节点信息的指针.
核心操作是 union 和 find.
* find: 找到根节点
    * 递归向上找到根节点并返回
* union: 将两个元素合并到一个集合中
    * x,y 两个元素在一个树中, 直接返回
    * 不再一棵树, 随便吧一个节点加入另一个节点所在的树中
