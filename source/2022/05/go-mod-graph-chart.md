---
git: "https://github.com/PaulXu-cn/go-mod-graph-chart"
tags:
- 源码
- golang
- graph
- dependency
- 依赖管理
date: "2022-05-30"
---

# go-mod-graph-chart 源码阅读

[go-mod-graph-chart](https://github.com/PaulXu-cn/go-mod-graph-chart)

## 简介

一个分析 `go mod graph` 的工具

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Go                           7     30686       47        27    30612        109
JSON                         5     10690        0         0    10690          0
JavaScript                   5       583       95        73      415         30
Markdown                     2        75       23         0       52          0
CSS                          1        65        8         1       56          0
HTML                         1        14        0         0       14          0
License                      1        21        4         0       17          0
gitignore                    1        21        4         4       13          0
───────────────────────────────────────────────────────────────────────────────
Total                       23     42155      181       105    41869        139
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $1,363,272
Estimated Schedule Effort 15.472319 months
Estimated People Required 7.827860
───────────────────────────────────────────────────────────────────────────────
Processed 1733808 bytes, 1.734 megabytes (SI)
───────────────────────────────────────────────────────────────────────────────
```

特别声明大部分代码可能是生成的, 所以核心功能代码量不大.

## 源码

看的目的是作者如何解决 `go mod graph` 中存在的一些问题

1. 依赖会存在循环
    1. 不控制会导致依赖树变为一个有向循环图
2. 太多如何打印
    1. 不控制会导致打印太多,有一个项目我打印了 1.7G 的文件(主要不知道什么时候结束, 我主动停了程序, 实际是多少不知道)
    2. `golang/x` 这这种包太混乱了, 没有明确的版本导致依赖混乱至极

这两个问题是我之前写 [dependency-graph](https://github.com/zoroqi/dependency-graph) 遇到的. 我的解决方案时使用了特殊标志位来控制, 一个代表出现了循环, 一个代表曾经打印过

### 依赖循环

看代码是没有处理这两种情况, 但输出结果和我的不一致, 测试项目 [sanke](https://github.com/1024casts/snake)

`mod graph` 结果中 `github.com/swaggo/gin-swagger` 依赖了 13 个, 而本项目只依赖了 10 个, 

#### 主要代码

核心 tree 结构
```go
type Tree struct {
	Id       uint32  `json:"-"`
	Name     string  `json:"name"`
	Value    uint32  `json:"value"`
	Children []*Tree `json:"children"`
}
```

辅助关系构建结构
```go
type RouteNode struct {  
   Id    uint32   `json:"id"`
   Name  string   `json:"name"`
   Value uint32   `json:"value"`
   Route []uint32 `json:"-"`
}
```

索引用的三个字典进行存储
1. 全局的 RouteNode, key: string, value: RouteNode
2. 返回结果的 repeatDependNodes, key: uint32, value: `*Tree`
3. 方法中定义的 roots, key: string, value:`*Tree`

在构建 `*Tree` 至少从名字上感觉是反着的, 读起来很别扭.
```go
newRepeatDependNode = Tree{
   Name: tree2.Name,
   Value: tree2.Value,
   Id: tree2.Id,
   Children: []*Tree{
      {
         Name: theParentNode.Name,
         Value: theParentNode.Value,
         Id: theParentNode.Id,
      },  
      {  
         Name: theRouteNode1.Name,
         Value: theRouteNode1.Value,
         Id: theRouteNode1.Id,
      },
   },
}
```

最后, 返回结果 repeatDependNodes 中的 Tree.Children 记录的是被依赖包

而内部定义 roots 中的 Tree.Children 记录的是依赖包, 而返回结果中的 root 是一个依赖树

全局定义的 RouteNode 提供了两者的关联, 不知道为啥要这样做. RouteNode 中 Route 记录依赖包 Tree.Children 数组下标, 感觉怪怪的.

## 简单说说

代码没看太懂, 反正这两个结果搞得好麻烦, 还不如直接 `map[string]*Tree` 可以不用记录下标了. 我通常构建树的时候是不会选择使用数组这种结构, 数组下标操作太费劲了. 真的是太费劲了, 为什么这么做就不知道了. 

> 补充: 看了下我的代码, 用的也是数组. 我构建树的方式使用的栈来进行遍历的; 我最开始应该是用的递归实现的, 因为测试项目发现对`mod graph`解析发现存在环, 直接栈溢出, 我迫不得已换成了栈.

构建被依赖关系, 我实现倾向与先构建好依赖树, 之后通过遍历依赖树生成相应的被依赖关系, 这样做会简单很多. 至少两个逻辑是分离的, 没有耦合在一起.

最后我要找的问题没有找到, 而且感觉有 bug.
