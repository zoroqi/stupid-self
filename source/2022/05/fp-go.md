---
git: "https://github.com/repeale/fp-go"
tags:
- 源码
- golang
- fp
- 函数式
- 流式编程
date: "2022-05-17"
---

# fp-go 源码阅读

[fp-go](https://github.com/repeale/fp-go)

## 简介

一个简单的 golang 库, 可以让用户使用类似函数式中的 map, filter 进行处理.

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Go                          23       900      184        39      677        128
XML                          3        20        0         0       20          0
gitignore                    2        23        3         8       12          0
License                      1        21        4         0       17          0
Markdown                     1       245       85         0      160          0
YAML                         1        27        5         0       22          0
───────────────────────────────────────────────────────────────────────────────
Total                       31      1236      281        47      908        128
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $24,410
Estimated Schedule Effort 3.355031 months
Estimated People Required 0.646405
───────────────────────────────────────────────────────────────────────────────
Processed 31565 bytes, 0.032 megabytes (SI)
───────────────────────────────────────────────────────────────────────────────
```


## 源码

整体评价对 1.8 泛型的快速应用. 针对流式的代码, 我更倾向于使用 chan 而不是用 slice. 我的原因比较简单, 当流程比较长且原始的数据比较多的时候 chan 更节省内存, 只需要最后跟一个 ToSlice 函数就可以提取所有结果, 也很方便.

### filter
```go
func Filter[T any](predicate func(T) bool) func([]T) []T
```

### map
```
func Map[T any,R any](predicate func(T) R) func([]T) []R
```

### pipe 和 compose

这两个是相似的, 只是方向不同

```
func Compose2[T1, T2, R any](fn1 func(T2) R, fn2 func(T1) T2) func(T1) R
func Pipe2[T1, T2, R any](fn1 func(T1) T2, fn2 func(T2) R) func(T1) R
```

- pipe 是 `fn2(fn1(t))`
- compose 是 `fn1(fn2(t))`

### flat

类似 haskell 的 concat 函数, 列表拆开合并, 但 flat 输入只能是二维 slice

### curry

提供一个柯里化的函数, 这就出现了复杂函数声明语句
```go
func Curry4[T1, T2, T3, T4, R any](fn func(T1, T2, T3, T4) R) func(T1) func(T2) func(T3) func(T4) R {
```

伟大的返回结果`func(T1) func(T2) func(T3) func(T4) R` 这是一个类型声明

### every 和 some

* every 等于 heakell 的 all 函数
* some 等于 heakell 的 any 函数

### reduce

对数据进行 map 合并到 acc 变量上. 可以理解为 foldl

```go
func Reduce[T any, R any](callback func(R, T) R, acc R) func([]T) R
```

## 一些吐槽

不知道为啥在 filter, map 等函数会有两个特殊的参数版本, 一个是提供元素下标, 一个提供下标和原始 slice 的, 不知道具体的应用场景是啥.

```golang
predicate func(T, int) bool
predicate func(T, int, []T) bool
```

## golang 的坑

* [[03/go泛型尝鲜]]
* [[22-03-16-21|go的泛型也是没谁了]]

我进行尝试关于流的处理, 我倾向于定义一个结构, 使用 chan 进行沟通. 这里面遇到一个问题就是不支持 map 操作. 之后在在 [函数式编程在 go 泛型下的实用性探索 — 银色子弹](https://silverrainz.me/blog/funtional-programming-in-go-generics.html) 这片文章中提到了部分解释.

> 提案的 [No parameterized methods](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods) 一节明确表示了方法（method，也就是有 recevier 的函数）不支持单独指定类型参数
> 
> ![[202205122134#链式调用]]

我心中最完美的方案是这样的, 但已经不可能了.
```golang
type Stream[T any] struct {
    c chan T
}
func (s *Stream[T]) Filter(p func(T) bool) Stream[T]
func (s *Stream[T]) Map[T2 any](m func(T)T2) Stream[T2]
```

还有就是 golang 没有语法糖, 导致编写这种 lambda 表达式很费劲, 稍微写长点就导致缩进和嵌套很混乱. 但其实也好, lambda 使用过多容易导致代码不好维护.
