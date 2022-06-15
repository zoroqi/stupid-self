---
github: "https://github.com/golang/groupcache"
tags:
- 源码
- golang
- 缓存
- google
date: "2022-06-14"
---

# groupcache 源码阅读

[groupcache](https://github.com/golang/groupcache)

## 简介

一个简单的 golang 库, 可以让用户使用类似函数式中的 map, filter 进行处理.

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Go                          16      2946      354       541     2051        476
XML                          3        25        0         0       25          0
Protocol Buffers             2        97       14        30       53          0
gitignore                    2         9        0         3        6          0
License                      1       191       36         0      155          0
Markdown                     1        74       24         0       50          0
YAML                         1        19        4         0       15          0
───────────────────────────────────────────────────────────────────────────────
Total                       26      3361      432       574     2355        476
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $66,402
Estimated Schedule Effort 4.907323 months
Estimated People Required 1.202146
───────────────────────────────────────────────────────────────────────────────
Processed 91555 bytes, 0.092 megabytes (SI)
───────────────────────────────────────────────────────────────────────────────
```


发现是 golang 的官方库, 但看提交记录已经很久没有更新了. 

最让我惊讶的是竟然没有 go.mod 文件. 这是已经放弃了这个项目吗? 就不考虑把项目搞得规范一点吗?

## 源码

### cache 设计

核心结构
```go
type Group struct {  
   name       string
   getter     Getter // 缓存加载逻辑
   peersOnce  sync.Once
   peers      PeerPicker
   cacheBytes int64
   mainCache cache
   hotCache cache
   loadGroup flightGroup
   _ int32 // 内存对齐
   Stats Stats // 统计信息
}
```

获取数据流程
1. 从 mainCache 获取数据, 没有从 hotCache 获取数据
2. 还是没有使用 loadGroup 进行数据加载
    1. 尝试用 peers 获取元素
        1. 填充到 hotCache 中
        2. 之填充 10% 的数据
    2. 调用 getter 获取元素
        1. 填充到 mainCache 中
        2. 根据内存使用是情况, 决定要不要移除一部分元素

peers 是一个远程获取的方式获得 value 方法.

mainCache 缓存自身数据, hotCache 缓存其他节点的信息, hotCache 主要是防止热点数据不断从其他节点拉去.

### lru 算法

本地 cache 结构是采用链表设计 lru 算法实现.

核心结构
```go
type Cache struct {
    MaxEntries int 
    OnEvicted func(key Key, value interface{})
    ll    *list.List // 链表是双向链表, 这里保存链表首元素
    cache map[interface{}]*list.Element
}
```

特别声明, 整个 Cache 是不保证线程安全的, 需要外部进行安全保证. 

1. 添加操作
    1. element 写入 cache 中
    2. 将 element 移动到链表头部
    3. 判断是否超出上线
        1. 超了执行最久未使用元素
2. 移除操作
    1. 从链表中移除 element
    2. 将 element 从 cache 中 delete 
3. 超出最久没有使用元素
    1. 从链表中获取尾元素
    2. 执行移除操作
4. 获取元素
    1. 从 cache 中查找元素 element
    2. 将 element 移动到链表头部

## 一些吐槽

### 一个无法理解的设计
函数也可有方法, 虽然不知道这么做在设计上有什么好处, 但很震惊. 这么做以为这 GetterFunc 还可以实现一些接口.
```go
type GetterFunc func(key string) (interface{}, bool)
func (f GetterFunc) Get(key string) (interface{}, bool) {
    return f(key)
}
```

测试编译感觉可以实现一种特殊的接口嵌套, 一下写法就可以实现 Cache 没有显性实现 Getter 接口, 而是通过嵌套 GetterFunc 获得了实现.
```go
type Getter interface {  
   Get(string) (interface{}, bool)  
}  
  
type Cache struct {  
   GetterFunc  
}
```

### 缩写
golang 建议使用缩写或简写来来对变量命名, 但是看起来有点费劲. 

我对缩写和简写的观点倾向于和变量的使用范围有关, 使用范围越广就不要简写了, 写清楚最好; 作用范围小的就随意了, 只要不产生误导就好了, 比如"io.Reader 变量起名 w"就是一个反面例子, 简写最好叫"r".

### 一种全局设计

为了保持单例, 使用了一个全局变量缓存很多东西, 包括缓存组, Peer 初始化方式, 等操作都是通过一下方式处理的
```
// 注册某些东西
Register....(fn func()) 
```

在内部根据需要进行初始化, 这中处理就感觉怪怪的, 当然看代码风格是为了控制代码. 

感觉通过全局方法进行操作可以在基于 [[gossip]] 协议动态调整节点信息. 反正我感觉怪怪的, 我不是很喜欢这种存在全局状态的信息.