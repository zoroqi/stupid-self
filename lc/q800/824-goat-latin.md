---
aliases:
- 824. 山羊拉丁文
- 824. goat latin
tc:
- leetcode
- algorithm
leetcode:
  num: 842
  url: https://leetcode.cn/problems/goat-latin/
  tags:
  - 字符串
id: 20221012215328_d91116d700bc4428
date: "2022-10-12"
---

# 824. 山羊拉丁文

给你一个由若干单词组成的句子 `sentence` ，单词间由空格分隔。每个单词仅由大写和小写英文字母组成。

请你将句子转换为 *“*山羊拉丁文（*Goat Latin*）*”*（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。山羊拉丁文的规则如下：

*   如果单词以元音开头（`'a'`, `'e'`, `'i'`, `'o'`, `'u'`），在单词后添加`"ma"`。
    *   例如，单词 `"apple"` 变为 `"applema"` 。
*   如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加`"ma"`。
    *   例如，单词 `"goat"` 变为 `"oatgma"` 。
*   根据单词在句子中的索引，在单词最后添加与索引相同数量的字母`'a'`，索引从 `1` 开始。
    *   例如，在第一个单词后添加 `"a"` ，在第二个单词后添加 `"aa"` ，以此类推。

返回将 `sentence` 转换为山羊拉丁文后的句子。

```
示例 1：

输入：sentence = "I speak Goat Latin"
输出："Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
示例 2：

输入：sentence = "The quick brown fox jumped over the lazy dog"
输出："heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
```

**提示：**

* `1 <= sentence.length <= 150`
* `sentence` 由英文字母和空格组成
* `sentence` 不含前导或尾随空格
* `sentence` 中的所有单词由单个空格分隔

## 愚笨的思考

很有意思的一个规则, 也很好玩, 简单写一个看看.

用 haskell 的思路是将字符串拆分成词组, 和一个数字进行组合成为元组, 对元组进行转换处理, 最后在转会字符串.

写的过程中遇到一个问题, `case of` 语法不成功, 之前一直以为是这个语法和函数的匹配没有什么区别, 我就写的是
```hs
case x y of
1 2 -> 3
3 4 -> 7
n b -> n - b
```
而这种方式会对下面的匹配产生报错, 意向中看的教程是写的 case 表达式 of, 看来只能写一个不能匹配多个, 需要把写成 `case (x,y) of` 就可以了. 看官方教程 [A Gentle Introduction to Haskell: Patterns](https://www.haskell.org/tutorial/patterns.html) 中 4.3  Case Expressions 说明了, 需要进行元组的包装.
