---
github: https://github.com/yuin/goldmark-meta
tags:
- 源码
- markdown
- yaml
- golang
date: "2022-05-17"
id: 20220517155609_b02f07dc66f344be
---

# goldmark-meta 源码阅读

[goldmark-meta](https://github.com/yuin/goldmark-meta)

## 简介

[yuin/goldmark](https://github.com/yuin/goldmark) 的一个插件, 主要是针对 markdown 文件中起始 meta 信息的解析.

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
XML                          3        20        0         0       20          0
Go                           2       593       59        22      512         88
gitignore                    2        21        2         6       13          0
License                      1        21        4         0       17          0
Markdown                     1       189       34         0      155          0
───────────────────────────────────────────────────────────────────────────────
Total                        9       844       99        28      717         88
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $19,049
Estimated Schedule Effort 3.053321 months
Estimated People Required 0.554286
───────────────────────────────────────────────────────────────────────────────
Processed 18209 bytes, 0.018 megabytes (SI)
───────────────────────────────────────────────────────────────────────────────

App elapsed:  7.919094ms
```

是 goldmark 的插件所以代码量很少, 写了一个批量插入 meta 信息的工具, 所以看看是咋实现这段逻辑解析的.

## 解析 meta

我实现的方式比较暴力, 就是行读取然后从 0 行开始处理到第二个"`---`"之间的部分.

看了代码后发现这个插件逻辑是一样的, 当然他依赖于 goldmark 的文法解析器进行处理. 把 md 的 meta 部分当做一个特殊的 block 进行的处理, 只是这个 block 不是用反引号或波浪线包裹, 而是用的短线进行包裹.

## 新发现
### yaml 组件
我解析 yaml 用的组件和作者一样都是 [go-yaml/yaml](https://github.com/go-yaml/yaml). 

我用的时候遇到一个问题"在解析 yaml 文件必须用 map 承接, 导致输出的时候顺序和读入就不一致了". 对我来说影响不大, 因为顺序变化对程序来说没有什么大影响, 就人主观有一点不爽; 作为个人工具也就没着急处理. 看这个项目突然发现 yaml 有一个 MapSlice 的类型可以解决这个问题, 后边升级我的插件吧.

### 一个循环

没有这么实现过循环, 但从效果来说是初始化一个集合并赋值. 我可能会使用 `_,_` 来做写, 然后 IDE 提示再删掉.

```go
var arr []MyInt
for range list {  
   arr = append(arr, MyInt(10))  
}
```

### go import 中的横线

项目名称叫 "goldmark-meta" 中间有一个横线, 很好会咋处理, 发现这个时候 golang 在语法中不能定义 `m := goldmark-meta.Meta` 的东西. 最后可以直接用横线后半部分作为包名 `meta.Meta`, 也是神奇. 

当然我建议还是 import 的时候起一个名字比较好.

## 一些无法理解的东西

看过几个 golang 的项目, 发现很喜欢使用一下方式进行值设置.
```
func New(opts ... Option) Server {
    // ....
}
```

这个项目也是采用这个方式进行配置设置, 不是很懂这么做事为了什么. 看实现来说是可以自己实现一个 Option 的, 但想要修改的值通常还是内部属性, 自己实现也是没有意义的. 代码最后有啥用呢?
