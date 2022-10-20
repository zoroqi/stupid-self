---
aliases:
- 802. 找到最终的安全状态
- 802. find eventual safe states
tc:
- leetcode
- algorithm
leetcode:
  num: 802
  url: https://leetcode.cn/problems/find-eventual-safe-states/
  tags:
  - 图
  - 拓扑
  - 深度优先遍历
  - 广度优先遍历
date: "2022-10-20"
id: 20221020181631_4f2a8a325acf4964
---

# 802. 找到最终的安全状态

有一个有 n 个节点的有向图，节点按 0 到 n - 1 编号。图由一个索引从 0 开始的 2D 整数数组 graph 表示， graph[i]是与节点 i 相邻的节点的整数数组，这意味着从节点 i 到 graph[i]中的每个节点都有一条边。

如果一个节点没有连出的有向边，则它是终端节点 。如果没有出边，则节点为终端节点。如果从该节点开始的所有可能路径都通向终端节点 ，则该节点为安全节点 。

返回一个由图中所有 **安全节点** 组成的数组作为答案。答案数组中的元素应当按 升序 排列。

```
示例 2：
输入：graph = [[1,2],[2,3],[5],[0],[5],[],[]]
输出：[2,4,5,6]
解释：示意图如上。
节点 5 和节点 6 是终端节点，因为它们都没有出边。
从节点 2、4、5 和 6 开始的所有路径都指向节点 5 或 6 。

示例 2：

输入：graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]
输出：[4]
解释:
只有节点 4 是终端节点，从节点 4 开始的所有路径都通向节点 4 。
```

**提示：**

* `n == graph.length`
* `1 <= n <= 10^4`
* `0 <= graph[i].length <= n`
* `0 <= graph[i][j] <= n - 1`
* `graph[i] 按严格递增顺序排列。`
* `图中可能包含自环。`
* `图中边的数目在范围 [1, 4 * 10^4] 内。`

## 愚笨的思考

根据示例, 我感觉满足一下任意条件的节点就是安全节点
1. 没有出度
2. 所有出度指向没有出度节点(终结节点)

另一种判断安全节点判断条件
1. 没有出度
2. 节点不能在某个环上
3. 节点不能指向某个环

这里的歧义是我不是很理解的语句, "所有可能路径" 是指临近的, 还是说只要可以跳转到就可以, 不是很理解这个"路径"的表述.
并且给出的示例无法区分上下这两种情况.

感觉我理解的"路径"我倾向于下边的条件, 实现根据一下条件实现.


### 暴力方案 PlanA

对深度优先遍历, 把能成环和进入环上的节点删除, 就可以了. 剩下的就是安全的节点.

直接暴力果然会超时, 蛋疼啊.
做了一下几个小优化, 将第 98 (节点数 500 个) 测试用例, 的递归调用次数减少的数量.
1. 对环中的节点一并加入到不安全.
2. 判断环的时候, 一并判断不安全节点集合.
3. 我需要指向安全的节点就不用再判断环了.

优化过程中发现我好想掉沟里了, 我任何实现了一个单独用来判断环的函数, 这里不进行全局状态的修改, 我是可以在这里进行修改状态, 来减少遍历的.
用时 440 ms,
逐渐删除重复逻辑最后耗时在 90 ms 左右

haskell 完全不知道如何实现.
