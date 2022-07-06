---
aliases:
- 844. 比较含退格的字符串
- 844. backspace string compare
tc:
- leetcode
- algorithm
leetcode:
    num: 844
    url: https://leetcode.cn/problems/backspace-string-compare/
    tags:
    - 字符串
    - 双指针
    - 栈
---

# 844. 比较含退格的字符串

给定 `s` 和 `t` 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 `true` 。`#` 代表退格字符。

**注意：**如果对空文本输入退格字符，文本继续为空。

```
示例 1：

输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。
示例 2：

输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。
示例 3：

输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。
```

**提示：**

* `1 <= s.length, t.length <= 200`
* `s` 和 `t` 只含有小写字母以及字符 `'#'`

**进阶：**

* 你可以用 `O(n)` 的时间复杂度和 `O(1)` 的空间复杂度解决该问题吗？

## 愚笨的思考

### 失败方案

第一反应, 正则可能能搞定, `.#` 替换成空, 就可以了. 只要保证没有连续的 "#" 就可以了, 但用例发现可以有连续的 "#". 看来正则是不行了

### 暴力方案 PlanA

生成对应的结果就好了, 使用栈来处理, 遇到"#"直接出栈, 最后完全出栈比较就可以了. golang 比较简单.