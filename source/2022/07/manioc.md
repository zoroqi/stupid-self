---
github: https://github.com/fuzmish/manioc
tags:
- 源码
- golang
- IOC
- 泛型
date: "2022-06-14"
tc:
- github
- reading_code
id: 20220614094736_0326b08aede7491c
---

# manioc 源码阅读

[manioc](https://github.com/fuzmish/manioc)

## 简介

一个简单的 golang 库, 提供依赖注入功能. 代码还在 alpha 阶段, 是不能用在生产的, 很多东西还没有确定下来.

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Go                          24      2924      521       342     2061        107
XML                          3        25        0         0       25          0
JSON                         2        22        0         0       22          0
SVG                          2       898        0         0      898          1
YAML                         2        49        7         0       42          0
gitignore                    2        10        0         3        7          0
Dockerfile                   1        18        1         4       13          9
License                      1        21        4         0       17          0
Makefile                     1        11        2         0        9          1
Markdown                     1       344       60         0      284          0
───────────────────────────────────────────────────────────────────────────────
Total                       39      4322      595       349     3378        118
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $96,981
Estimated Schedule Effort 5.667030 months
Estimated People Required 1.520367
───────────────────────────────────────────────────────────────────────────────
Processed 476263 bytes, 0.476 megabytes (SI)
───────────────────────────────────────────────────────────────────────────────
```

## 源码

主要看的是如何处理依赖的

### 简单使用和注意事项
```go
type SiA interface {  
   Do()  
}  
type SiB interface {  
   Run()  
}  
  
type SeA struct {}  
  
func (s *SeA) Do() {  
   fmt.Println("sea")  
}  
  
type SeB struct {
   a SiA  `manioc:"inject"`
   //c SiC `manioc:"inject"`
}  
  
func (s *SeB) Run() {  
   s.a.Do()  
   fmt.Println("seb")  
}

type SiC interface {  
   Run()  
}  
  
type SeC struct {  
   b SiB `manioc:"inject"`  
}  
  
func (s *SeC) Run() {  
   s.b.Run()  
   fmt.Println("sec")  
}

func main() {
    manioc.Register[SiA, SeA]()  
    manioc.Register[Sib, SeB]()
    manioc.Register[SiC, SeC]()
    c,_ := manioc.Resolve[SiC]()
    c.Run() 
}
// 输出
// sea
// seb
// sec
```

### 加载过程

加载根据缓存策略不同进行了不同的加载方式, 默认每次初始化一个新的, 这种初始化会级联把所有的的都进行初始化.

使用一种类似执行链的方式进行初始化 struct. 效果类似下面这种效果

```go
// 创建一个 struct
type activator func(ctx resolveContext) (any, error)

func newStruct(ctx resolveContext) (any, error) {
    return new(...),nil
}

func inject(ctx resolveContext,init activator)(any, error) {
    return func(ctx resolveContext) (any, error)
        instance, err := init()
        if err != nil {
            return nil,err
        }
        // 处理 inject 逻辑
        return instance, nil
    }
}
func cachePolicy(ctx resolveContext,init activator)(any, error) {
    return func(ctx resolveContext) (any, error)
        instance, err := init()
        if err != nil {
            return nil,err
        }
        // 处理缓存策略逻辑
        return instance, nil
    }
}
```

通过嵌套组合出想要的结果, 比如不需要缓存那就 `newStruct 和 inject `进行组合. 最终将这个调用链缓存在一个 register 结构是 `map[key][]activator`, 对获得的数据 activator 直接调用结果进行一次强转然后返回结果就好了.

所以看代码并没有解决循环依赖的问题, 测试结果也是直接爆栈引起系统崩溃.

## 一些吐槽

### golang 的 struct 的优化

```
type A struct{}
type B struct{}
func main() {
    a1 := new(A)
    a2 := new(A)
    b1 := new(B)
    fmt.Println(a1 == a2) //true
    fmt.Println(a1 == b1) //true
}
```

原来 `struct{}` 被进行过特殊优化, 以前没有注意过, 都指向相同的抵制. 很有意思

### 比 java 泛型好用一些

看代码发现 golang 的泛型要比 java 好用很多, 比如 golang 是可以初始化具体的结构的, 但 java 是不行的. 

比如:
```go
type A struct{}
func NewPoint[T any]() *T {  
   return new(T)  
}
```

java 是不能这么做的, 无法初始化一个数据, 所以部分 java 库为了解决这问题, 通过继承一个类, 在这个内部类上写明具体类型, 通过反射就可以拿到了, 就可以初始化了.

golang 就不需要了, 直接根据需要创建, 但是这种方式是无法对字段进行赋值的, 而这也就一样需要通过 #反射 处理了.

golang 的泛型提供的是一种 interface 的处理, 所以无法明确说明字段, 但也可以通过 setter 方式来控制
