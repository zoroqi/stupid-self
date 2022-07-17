---
aliases:
- 846. 一手顺子
- 846. hand-of-straights
tc:
- leetcode
- algorithm
leetcode:
    num: 846
    url: https://leetcode.cn/problems/hand-of-straights/
    tags:
    - 贪心
    - 数组
    - 哈希表
    - 排序
---

# 844. 比较含退格的字符串

Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，使每一组的牌数都是 `groupSize` ，并且由 `groupSize` 张连续的牌组成。

给你一个整数数组 `hand` 其中 `hand[i]` 是写在第 `i` 张牌，和一个整数 `groupSize` 。如果她可能重新排列这些牌，返回 `true` ；否则，返回 `false` 。

```
示例 1：

输入：hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
输出：true
解释：Alice 手中的牌可以被重新排列为 [1,2,3]，[2,3,4]，[6,7,8]。
示例 2：

输入：hand = [1,2,3,4,5], groupSize = 4
输出：false
解释：Alice 手中的牌无法被重新排列成几个大小为 4 的组。
```

**提示：**

* `1 <= hand.length <= 10^4`
* `0 <= hand[i] <= 10^9`
* `1 <= groupSize <= hand.length`

## 愚笨的思考

第一个问题是题没有完全看懂, 是一定保持均分吗? 看示例是均分的. 连续是表示自增的意思

### 直接统计 PlanA

连续即隐含了一下的描述: 假设 `groupSize = 4` 那一定存在连续的 `n1,n2,n3,n4`, 的统计数量保证 `n1.size <= n2.size <= n3.size <= n4.size`

算法逻辑就是
1. 统计每一个数字出现的次数
2. 对统计后的数字进行排序
3. 遍历有序数组
    1. 进行 $n_0$ 到 $n_groupSize$ 范围遍历
        1. 从统计中 $n_i$ 减去的 $n_0$ 的统计量
        2. 修改 $n_i$ 中的统计量
        2.  $n_i$ 小于 0 则返回 false, 元数据无法进行重组
4. 直到循环结束返回 true

排序是必须的, 不能在不排序情况下进行处理, 因为无法保证 "n1" 后边一定存在后续数字

golang 耗时直接是 40ms 这就达到了 85% 的水平, 我以为 40ms 是垫底的. 用第一名的代码, 我发现耗时也是 40ms, 看来数据不准了.

haskell 怎么实现就比较麻烦了, golang 可以简单的循环和修改变量 haskell 没有.

我可以用 group 进行合并, 根据后一个数组是否可以进行 take 出相同的数据来判断, 但这些起来真的很费劲, 算了.
