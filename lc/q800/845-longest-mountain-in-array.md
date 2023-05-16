---
aliases:
- 845. 数组中的最长山脉
- 845. longest mountain in array
tc:
- leetcode
- algorithm
leetcode:
  num: 845
  url: https://leetcode.cn/problems/longest-mountain-in-array
  tags:
  - 数组
  - 双指针
date: "2023-02-16"
id: 20230216041009_d75c1fa83c0f4a40
---

# 845. 数组中的最长山脉

把符合下列属性的数组 `arr` 称为 **山脉数组** ：

* `arr.length >= 3`
* 存在下标 `i`（`0 < i < arr.length - 1`），满足
    * `arr[0] < arr[1] < ... < arr[i - 1] < arr[i]`
    * `arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`

给出一个整数数组 `arr`，返回最长山脉子数组的长度。如果不存在山脉子数组，返回 `0` 。

```
示例 1：

输入：arr = [2,1,4,7,3,2,5]
输出：5
解释：最长的山脉子数组是 [1,4,7,3,2]，长度为 5。
示例 2：

输入：arr = [2,2,2]
输出：0
解释：不存在山脉子数组。
```

**提示：**

* `1 <= arr.length <= 10^4`
* `0 <= arr[i] <= 10^4`

**进阶：**

* 你可以仅用一趟扫描解决此问题吗？
* 你可以用 `O(1)` 空间解决此问题吗？

## 愚笨的思考


找到拐点数据, 并记录左右最低点的索引位置就好了.
根据题目描述递增, 递减, 先递减再递增, 水平都不构成"山脉".

## 暴力方案 PlanA

两个指针记录起始和结束索引, 在出现拐点的时候进行记录.

~~记录所有的向上的拐点信息, 对两个拐点获得差值, 最大的差值就是结果.~~

这里需要记录两个值, 需要记录山的起始点, 也需要记录结束点. 山脉不能出现"水平"的情况.

三个元素总共可以构成 9 种状态, 分别是
1. 山, 坑
2. 水平, 上坡, 下坡
3. 上弯, 下弯
3. 折上, 折下

生成一个简单的方向性数组, 有三种状态, 和前一个比的方向就好了.

但是感觉并不是很好实现, 需要太多的判断了, 存在多处需要特殊处理的地方, 还是算了吧.

## 找到山顶后找山脚 PlanB

进行一次遍历找到所有山顶, 并通过山顶反向遍历找到两个山脚再计算距离, 这样明显简单一些.
这个实现的逻辑比较简单, 就是对整个数组进行两次遍历, 当然也可以压缩到一个大循环里面. 那样代码就有点复杂了, 我还是倾向于这么写.

但不知道如何用 haskell 实现这个逻辑. 之后看看如何实现吧.