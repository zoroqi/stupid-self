---
aliases:
- 884. 两句话中的不常见单词
- 884. uncommon words from two sentences
tc:
- leetcode
- algorithm
leetcode:
  num: 884
  url: https://leetcode.cn/problems/uncommon-words-from-two-sentences/
  tags:
  - 哈希
  - 字符串
date: "2022-10-19"
id: 20221019163300_860c9102e22e4922
---

# 884. 两句话中的不常见单词

句子 是一串由空格分隔的单词。每个 单词 仅由小写字母组成。

如果某个单词在其中一个句子中恰好出现一次，在另一个句子中却 没有出现 ，那么这个单词就是 不常见的 。

给你两个 句子 s1 和 s2 ，返回所有 不常用单词 的列表。返回列表中单词可以按 任意顺序 组织。

```
示例 1：

输入：s1 = "this apple is sweet", s2 = "this apple is sour"
输出：["sweet","sour"]
示例 2：

输入：s1 = "apple apple", s2 = "banana"
输出：["banana"]
```

提示：

* `1 <= s1.length, s2.length <= 200`
* `s1 和 s2 由小写英文字母和空格组成`
* `s1 和 s2 都不含前导或尾随空格`
* `s1 和 s2 中的所有单词间均由单个空格分隔`

## 愚笨的思考

将字符串转换成词频统计, 过滤掉出现两次以上的词, 之后求两个集合差集的并集.

写的时候发现, 一个词只能出现一次, 那问题就变成了词频统计, 让后返回仅出现一次的集合就好了.

haskell 至少我这个不是很复杂, 网上看了几个发现, 词频统计的实现都很复杂
