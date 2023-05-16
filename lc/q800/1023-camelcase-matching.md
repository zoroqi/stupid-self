---
aliases:
- 1023. 驼峰式匹配
- 1023. camelcase matching
tc:
- leetcode
- algorithm
leetcode:
  num: 1023
  url: https://leetcode.cn/problems/camelcase-matching/
  tags:
  - 字符串
  - 字典树
  - 字符串匹配
date: "2023-05-16"
id: 20230516060905_509a5b94ad5149f3
---

# 1023. 驼峰式匹配

给你一个字符串数组 queries，和一个表示模式的字符串 pattern，请你返回一个布尔数组 answer 。只有在待查项 queries[i] 与模式串 pattern 匹配时， answer[i] 才为 true，否则为 false。

如果可以将小写字母插入模式串 pattern 得到待查询项 query，那么待查询项与给定模式串匹配。可以在任何位置插入每个字符，也可以不插入字符。

```
示例 1：

输入：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FB"
输出：[true,false,true,true,false]
示例：
"FooBar" 可以这样生成："F" + "oo" + "B" + "ar"。
"FootBall" 可以这样生成："F" + "oot" + "B" + "all".
"FrameBuffer" 可以这样生成："F" + "rame" + "B" + "uffer".
示例 2：

输入：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBa"
输出：[true,false,true,false,false]
解释：
"FooBar" 可以这样生成："Fo" + "o" + "Ba" + "r".
"FootBall" 可以这样生成："Fo" + "ot" + "Ba" + "ll".
示例 3：

输入：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBaT"
输出：[false,true,false,false,false]
解释： 
"FooBarTest" 可以这样生成："Fo" + "o" + "Ba" + "r" + "T" + "est".
```

提示：

- 1 <= pattern.length, queries.length <= 100
- 1 <= queries[i].length <= 100
- queries[i] 和 pattern 由英文字母组成

## 愚笨的思考

### 单纯的正则 PlanA

第一反应是应该可以用正则搞定, 所以尝试正则看看.

我理解的问题是 pattern 串以大写字母为分割, 在大写字母之间可以插入任意个小写字符(我理解是这个意思)
我可以把 "FB" -> `^F[a-z]*?B[a-z]*?$` 感觉是可以的, 试试看.

没有成功, 至少三个例子是对的.

失败的例子是 `["CompetitiveProgramming","CounterPick","ControlPanel"]` `"CooP"`, 结果是 `[false, false, true]`.
我可以理解成,**只要相对位置对就可以了**, 中间插入多少字符都不重要.

那可以把正则改改.

"FoBa" -> `^F[a-z]*o[a-z]*B[a-z]*a[a-z]*$` 应该就可以了, 效率就更烂了, 把 "?" 去掉吧, 不回溯了.

还是没有成功, 新的失败用例

`["aksvbjLiknuTzqon","ksvjLimflkpnTzqn","mmkasvjLiknTxzqn","ksvjLiurknTzzqbn","ksvsjLctikgnTzqn","knzsvzjLiknTszqn"]` `"ksvjLiknTzqn"` 结果 `[true, true, true, true, true, true]`

失败原因是我开头没有加入`[a-z]*`这部分.

这个实现好像很吃内存, 内存排名 5%, 用的有点多.

haskell 正则不知道如何使用, 决定让 AI 给一个例子, 然后完全不对, 还是要 google 有点烦躁
