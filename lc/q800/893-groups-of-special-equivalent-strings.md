`---
aliases:
- 893. 特殊等价字符串组
- 893. groups of special equivalent strings
tc:
- leetcode
- algorithm
leetcode:
  num: 893
  url: https://leetcode.cn/problems/groups-of-special-equivalent-strings/
  tags:
  - 数组
  - 字符串
  - 哈希表
date: "2023-05-19"
id: 20230519104156_cc045108ef3d4295
---

# 893. 特殊等价字符串组

给你一个字符串数组 words。

一步操作中，你可以交换字符串 words[i] 的任意两个偶数下标对应的字符或任意两个奇数下标对应的字符。

对两个字符串 words[i] 和 words[j] 而言，如果经过任意次数的操作，words[i] == words[j] ，那么这两个字符串是 特殊等价 的。

例如，words[i] = "zzxy" 和 words[j] = "xyzz" 是一对 特殊等价 字符串，因为可以按 "zzxy" -> "xzzy" -> "xyzz" 的操作路径使 words[i] == words[j] 。
现在规定，words 的 一组特殊等价字符串 就是 words 的一个同时满足下述条件的非空子集：

该组中的每一对字符串都是 特殊等价 的
该组字符串已经涵盖了该类别中的所有特殊等价字符串，容量达到理论上的最大值（也就是说，如果一个字符串不在该组中，那么这个字符串就 不会 与该组内任何字符串特殊等价）
返回 words 中 特殊等价字符串组 的数量。

```
示例 1：

输入：words = ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
输出：3
解释：
其中一组为 ["abcd", "cdab", "cbad"]，因为它们是成对的特殊等价字符串，且没有其他字符串与这些字符串特殊等价。
另外两组分别是 ["xyzz", "zzxy"] 和 ["zzyx"]。特别需要注意的是，"zzxy" 不与 "zzyx" 特殊等价。
示例 2：

输入：words = ["abc","acb","bac","bca","cab","cba"]
输出：3
解释：3 组 ["abc","cba"]，["acb","bca"]，["bac","cab"]
```

**提示：**

- `1 <= words.length <= 1000`
- `1 <= words[i].length <= 20`
- 所有 `words[i]` 都只由小写字母组成。
- 所有 `words[i]` 都具有相同的长度。

## 愚笨的思考

### 排序归一 PlanA

可以对字符进行**任意次数的交换**, 那我可以把所有内容归一后在排重就好了.
归一方式是对奇偶位上的字符分别进行排序就可以了.

拼接就是简单的中划线进行拼接就可以了.

直接成功, 简单暴力.

### 题解方案 PlanB

题解中的方法是统计各个字符的数量并进行排重, 这个性能更好一些.
我的实现在长度比较小的时候和题解没有太大的差别, 但是在长度过大的时候我的性能会有很大的衰退
`
