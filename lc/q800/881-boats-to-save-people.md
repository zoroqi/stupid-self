---
aliases:
- 881. 救生艇
- 881. boats to save people
tc:
- leetcode
- algorithm
leetcode:
  num: 881
  url: https://leetcode.cn/problems/boats-to-save-people/
  tags:
  - 贪心
  - 数组
  - 排序
date: "2023-05-17"
id: 20230517194042_9621021d592e468b
---

# 881. 救生艇

给定数组 people。people[i]表示第 i 个人的体重 ，船的数量不限，每艘船可以承载的最大重量为 limit。

每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。

返回 承载所有人所需的最小船数。

```
示例 1：

输入：people = [1,2], limit = 3
输出：1
解释：1 艘船载 (1, 2)
示例 2：

输入：people = [3,2,2,1], limit = 3
输出：3
解释：3 艘船分别载 (1, 2), (2) 和 (3)
示例 3：

输入：people = [3,5,3,4], limit = 5
输出：4
解释：4 艘船分别载 (3), (3), (4), (5)
```

提示：

- 1 <= people.length <= `5 * 10^4`
- 1 <= people[i] <= limit <= `3 * 10^4`

## 愚笨的思考

### 排序组合法 PlanA

想法是排序后最大和最小进行组合就可以了.

优先移重指针, 如何无法和小的进行组合, 就向左移动, 并对计数结果+1.
可以产生组合轻指针向右移动, 并对结果+1.
直到完成遍历为止

啊! 竟然成功了, 但是时间有点久, 用时 70ms 左右.

我只是直觉这个方法应该是对的, 但是我不知道如何证明. 感觉需要多去理解如何证明自己的算法是正确的.

根据 leetcode 的官方证明

> * 若他不能与体重最重的人同乘一艘船，那么体重最重的人无法与任何人同乘一艘船，此时应单独分配一艘船给体重最重的人。从 people 中去掉体重最重的人后，我们缩小了问题的规模，变成求解剩余 n−1 个人所需的最小船数，将其加一即为原问题的答案。
> * 若他能与体重最重的人同乘一艘船，那么他能与其余任何人同乘一艘船，为了尽可能地利用船的承载重量，选择与体重最重的人同乘一艘船是最优的。从 people 中去掉体重最轻和体重最重的人后，我们缩小了问题的规模，变成求解剩余 n−2 个人所需的最小船数，将其加一即为原问题的答案。
