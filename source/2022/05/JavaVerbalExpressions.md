---
github: https://github.com/VerbalExpressions/JavaVerbalExpressions
aliases:
- JavaVerbalExpressions 源码阅读 
tags:
- 源码
- java
- 正则
- 工具
id: 20221012214554_655ab0ac4f944e58
---

# JavaVerbalExpressions 源码阅读

[VerbalExpressions](https://github.com/VerbalExpressions/JavaVerbalExpressions)

## 简介

很好玩的一个项目, 通过口语化生成文本匹配, 将正则表达式进行封装.

代码量
```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Java                        10      2217      324       515     1378         38
XML                          2       215       19         1      195          0
License                      1        20        4         0       16          0
Markdown                     1       115       25         0       90          0
Shell                        1        10        1         1        8          1
YAML                         1        37        8         0       29          0
gitignore                    1         8        0         0        8          0
───────────────────────────────────────────────────────────────────────────────
Total                       17      2622      381       517     1724         39
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $47,858
Estimated Schedule Effort 4.333099 months
Estimated People Required 0.981243
───────────────────────────────────────────────────────────────────────────────
Processed 90374 bytes, 0.090 megabytes (SI)
───────────────────────────────────────────────────────────────────────────────
```

## 个人看法

自认为正则能力还是可以的, 通过这个方式写真的有点"脱了裤子放屁"的意思, 对我个人整体意义不大.

能对我产生意义的地方是提供正则的口语化注释. 

## 源码

代码核心结构
```
prefixes 开始 ^
source 正则部分
suffixes 当做栈使用, 控制分组的括号和 $
```

大部分都会使用 `(?:exp)` 会创建一个不分配组号的组.

剩下的大部分流程都是在不断的往 `source` 字段中拼对应的结构.

### 阅读中遇到的一些问题
Q: 为什么都要创建无组号分组?

A: 方便 `*, +, ?, {n,m}` 的控制, 不适用分组很有可能会无法正确标记前一个是什么

Q: 如何处理 group 的关闭括号的.

A: 通过 suffixes 作为栈使用, 创建分组的时候, 就入栈一个 `)` , 在关闭分组的时候, 直接判断是否包含`)`, 直接出栈最后一个元素.

既然是直接出栈, 看看能不能产生一个正则错误. 发现果然可以, 在关闭组之前使用 `endOfLine(拼接结束符$)` 直接导致出栈错误, 多拼接了一个`)`而没有拼接`$`, 正则错误. 这里建议对分组括号尽量使用一个队列处理.

当然`$`是可以出现在中间的, `$` 代表匹配字符串的结束, 出现在中间导致什么都无法匹配.

Q: 使用库可能出现的问题?

A: 描述正确是不会的, 但需要对自己描述过程想清楚.

比如
```
builder.capture().then("abc").count(2,5).endGr(); // ((?:abc){2,5})
builder.capture().then("abc").endGr().count(2,5); // ((?:abc)){2,5}
```

这两个, 进行匹配或验证是一样的, 但取分组是不一样的, 第一个可以获取到 `abcabc` 而第二个只能去到 `abc`
