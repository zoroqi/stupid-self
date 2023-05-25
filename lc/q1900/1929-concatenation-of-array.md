---
aliases:
- 1929. 数组串联
- 1929. concatenation of array
tc:
- leetcode
- algorithm
leetcode:
  num: 1912
  url: https://leetcode.cn/problems/concatenation-of-array/
date: "2023-05-25"
id: "20230525154221_a7d4f3b669b1450f"
---

# 1929. 数组串联

给你一个长度为 n 的整数数组 nums 。请你构建一个长度为 2n 的答案数组 ans ，数组下标 从 0 开始计数 ，对于所有 `0 <= i < n` 的 i ，满足下述所有要求：

ans[i] == nums[i]
ans[i + n] == nums[i]
具体而言，ans 由两个 nums 数组 串联 形成。

返回数组 ans 。


```
示例 1：

输入：nums = [1,2,1]
输出：[1,2,1,1,2,1]
解释：数组 ans 按下述方式形成：
- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]
- ans = [1,2,1,1,2,1]
示例 2：

输入：nums = [1,3,2,1]
输出：[1,3,2,1,1,3,2,1]
解释：数组 ans 按下述方式形成：
- ans = [nums[0],nums[1],nums[2],nums[3],nums[0],nums[1],nums[2],nums[3]]
- ans = [1,3,2,1,1,3,2,1]
```

**提示：**

提示：

提示：

- n == nums.length
- 1 <= n <= 1000
- 1 <= nums[i] <= 1000

## 愚笨的思考

看用例就是把数组复制一遍, 这个代码很简单. 想到了三种方式

1. 手动复制
2. 用 copy 函数
3. 直接 append 就好了

这里想到一个之前看到过的优化方式叫 `[[2022-09-01-11|达夫设备]]` 的优化技巧, 当然这是在编译上的一种优化技巧, 这里好像用不到.

haskell 可以直接 ++ ,也可以 `concat [a,a]`.
