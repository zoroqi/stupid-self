---
aliases:
- 833. 字符串中的查找与替换
- 838. find and replace in string
tc:
- leetcode
- algorithm
leetcode:
  num: 838
  url: https://leetcode.cn/problems/find-and-replace-in-string/
  tags:
  - 字符串
  - 数组
  - 排序
date: "2023-02-13"
---

# 833. 字符串中的查找与替换

你会得到一个字符串 `s` (索引从 0 开始)，你必须对它执行 `k` 个替换操作。替换操作以三个长度均为 `k` 的并行数组给出：`indices`, `sources`,  `targets`。

要完成第 `i` 个替换操作:

1. 检查 **子字符串**  `sources[i]` 是否出现在 **原字符串** `s` 的索引 `indices[i]` 处。
2. 如果没有出现， **什么也不做** 。
3. 如果出现，则用 `targets[i]` **替换** 该子字符串。

例如，如果 `s = "abcd"` ， `indices[i] = 0` , `sources[i] = "ab"`， `targets[i] = "eee"` ，那么替换的结果将是 `"eeecd"` 。

所有替换操作必须 **同时** 发生，这意味着替换操作不应该影响彼此的索引。测试用例保证元素间**不会重叠** 。

例如，一个 `s = "abc"` ，  `indices = [0,1]` ， `sources = ["ab"，"bc"]` 的测试用例将不会生成，因为 `"ab"` 和 `"bc"` 替换重叠。

*在对 `s` 执行所有替换操作后返回 **结果字符串** 。*

**子字符串** 是字符串中连续的字符序列。

```
输入：s = "abcd", indexes = [0,2], sources = ["a","cd"], targets = ["eee","ffff"]
输出："eeebffff"
解释：
"a" 从 s 中的索引 0 开始，所以它被替换为 "eee"。
"cd" 从 s 中的索引 2 开始，所以它被替换为 "ffff"。

输入：s = "abcd", indexes = [0,2], sources = ["ab","ec"], targets = ["eee","ffff"]
输出："eeecd"
解释：
"ab" 从 s 中的索引 0 开始，所以它被替换为 "eee"。
"ec" 没有从原始的 S 中的索引 2 开始，所以它没有被替换。
```

提示：

* `1 <= s.length <= 1000`
* `k == indices.length == sources.length == targets.length`
* `1 <= k <= 100`
* `0 <= indexes[i] < s.length`
* `1 <= sources[i].length, targets[i].length <= 50`
* `s` 仅由小写英文字母组成
* `sources[i]` 和 `targets[i]` 仅由小写英文字母组成


## 愚笨的思考


有两种方式, 拼接一个新的字符串, 或者从后面进行替换.
从后面进行替换可以保证前面的索引不变.

### 拼接字符串 PlanA

现实现拼接版, 假设 indices 是升序实现看看.
逻辑主要复杂的地方是需要处理中间的字符串, 进行不断的填充.

需要补充的内容有
1. 头元素
2. 两个 index 之间没有替换的元素
3. 尾部的元素

狗血的是没有排序好, 还要自己进行排序...

### 优化一下 PlanB

定一个 st 结构, 用来进行升序排序.
之后根据 PlanA 进行修改就好了.

测试效果还不错.

### 反向的替换 PlanC

对 st 进行降序排序.
这个实现比较简单, 直接截位, 替换, 替换.
测试数据保证了不会重叠.
