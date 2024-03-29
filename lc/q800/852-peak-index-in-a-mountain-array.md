---
aliases:
- 852. 山脉数组的峰顶索引
- 852. peak index in a mountain array
tc:
- leetcode
- algorithm
leetcode:
  num: 852
  url: https://leetcode.cn/problems/peak-index-in-a-mountain-array/
  tags:
  - 数组
date: "2023-02-23"
id: 20230223041027_7a141e1e87b14473
---

# 山脉数组的峰顶索引

符合下列属性的数组 `arr` 称为 **山脉数组** ：

* `arr.length >= 3`
* 存在 `i`（`0 < i< arr.length - 1`）使得：
    * `arr[0] < arr[1] < ... arr[i-1] < arr[i]`
    * `arr[i] > arr[i+1] > ... > arr[arr.length - 1]`

给你由整数组成的山脉数组 `arr` ，返回任何满足 `arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1]` 的下标 `i` 。

```
示例 1：

输入：arr = [0,1,0]
输出：1
示例 2：

输入：arr = [0,2,1,0]
输出：1
示例 3：

输入：arr = [0,10,5,2]
输出：1
示例 4：

输入：arr = [3,4,5,1]
输出：2
示例 5：

输入：arr = [24,69,100,99,79,78,67,36,26,19]
输出：2
```

**提示：**

* `3 <= arr.length <= 10^4`
* `0 <= arr[i] <= 10^6`
* 题目数据保证 `arr` 是一个山脉数组

**进阶：**很容易想到时间复杂度 `O(n)` 的解决方案，你可以设计一个 `O(log(n))` 的解决方案吗？

## 愚笨的思考

最开始以为是找到所有上峰的最大值, 后来看整个数组是一个山峰.

### PlanA

直接用的 [[845-longest-mountain-in-array|845. 数组中的最长山脉]] 中的判断逻辑就好了, 一次判决三个元素是否是山顶.

### PlanB

用 `O(log(n))` 的时间, 一定是二分查找应该可以.
问题中的描述是, 整个数组都符合山顶结构, 那就是二分查找找最大值.

实现上的逻辑就是判断现在中间元素是在左坡还是右坡.

但测试看性能并没有快多少, 和 PlanA 和 PlanC 在 leetcode 中的速度一样

这个代码不是我实现的, 第一次尝试用 Copilot 写代码, 感觉还不错.
虽然出了很久了, 但是一直没有使用, 开始作为生产力工具使用了.
测试看这对 CRUD 的代码有很好的效果

随着写 Haskell 的题增多, 现在写简单的循环转递归已经不用思考了, 反正靠着直觉写就好了, 尤其是在有现成的循环示例的话, 会更简单.

### PlanC

直接找最大值的索引, 这里需要舍弃首尾两个元素进行查找. 根据题目要求实际上, 直接找最大值也行.

想在 haskell 找一个类似的函数, 看看能不能直接用, 但是没有找到, 只能自己来了.
