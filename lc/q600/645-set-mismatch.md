---
aliases:
- 645. 错误的集合
- 645. set mismatch
tc:
- leetcode
- algorithm
leetcode:
  num: 645
  url: https://leetcode.cn/problems/set-mismatch/
  tags:
  - 数组
date: "2023-06-01"
id: "20230601141426_f70dc49c4ad04b3b"
---

# 645. 错误的集合

集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复。

给定一个数组 nums 代表了集合 S 发生错误后的结果。

请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

```
示例 1：

输入：nums = [1,2,2,4]
输出：[2,3]
示例 2：

输入：nums = [1,1]
输出：[1,2]
```

**提示：**

- `2 <= nums.length <= 10^4`
- `1 <= nums[i] <= 10^4`

## 愚笨的思考

### map 方式 PlanA

使用一个 map 进行统计, 之后对 map 进行遍历就可以了, 也可以使用 arr 进行统计.

### arr 方式 PlanB

逻辑和 map 一样, 只是进行统计的结构不同而已.

golang 在 `map[int]any` 和 slice 的操作语法完全一样, 可以无缝进行替换
