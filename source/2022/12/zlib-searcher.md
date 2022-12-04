---
id: 20221203203804_8d056b76174143e2
date: "2022-12-03"
aliases:
- zlib searcher 源码阅读
github: https://https://github.com/zu1k/zlib-searcher
tags:
- 源码
- rust
tc:
- github
- reading_code
---

# zlib searcher 源码阅读

## 简介

对 zlib 内容进行检索的一个辅助库. 通过扫描 csv 文件, 将书籍信息加载后提供对外检索能力.

```
─────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
─────────────────────────────────────────
Rust                         9       637       81        10      546         16
YAML                         7      1327      174         0     1153          0
TypeScript                   6        98        7         1       90          3
JSON                         5       124        0         0      124          0
TOML                         5       141       29         4      108          0
Vue                          5       419       22         1      396          4
gitignore                    4        44        4         7       33          0
Shell                        3       121       27         3       91         25
TypeScript Typings           3        19        4         3       12          0
XML                          3        20        0         0       20          0
Docker ignore                2        12        1         0       11          0
CSS                          1        10        1         0        9          0
Cargo Lock                   1      4945      477         2     4466          0
Dockerfile                   1        17        6         1       10          3
HTML                         1        12        0         0       12          0
License                      1        21        4         0       17          0
Makefile                     1        24        6         0       18          0
Markdown                     1       123       43         0       80          0
Python                       1        54        7         9       38          3
─────────────────────────────────────────
Total                       60      8168      893        41     7234         54
─────────────────────────────────────────
```

## 源码

本身不会 Rust, 对 Rust 相关的框架肯定是更不不知道了, 这个明显比 [[cdls]] 中的依赖组件要复杂.

特别说明, 没有搭建任何 Rust 的环境, 纯是看逻辑和猜测

这里用的 web 框架是一个叫 [actix/actix-web](https://github.com/actix/actix-web) 的框架.
可以 derive 来进行包装实现功能的快速扩展, 这个关键字想到了 haskell 的相同的关键字 deriving 了.
这个框架还可以提供类似 java 中逐渐的方式来实现配置, 比我用过的 gin 简单一些, 可能是我用 spring 习惯了导致的吧.

一下是一些有意思的代码片段

### 链式调用

```rust
let index_dir = std::env::current_exe()
    .unwrap()
    .parent()
    .unwrap()
    .join("index")
    .to_str()
    .unwrap()
    .to_string();
```

功能猜测是获取索引路径, 进行加载使用.
这段代码的 unwrap 让我想到之前看 "[haskell 使用 Either](https://lhbg-book.link/06-errors_and_files/01-either.html)" 中对连续处理的优化.
也是从类似连续 unwrap 转换成使用 `>=>` 来简化这个逻辑.
用 haskell 写可能写成这个效果(大概吧)
```haskell
index_dir = currentExe >=> parent >=> (join "index") >=> toStr >=> show
```

### 索引

索引使用 [quickwit-oss/tantivy](https://github.com/quickwit-oss/tantivy).
使用的 [messense/jieba-rs](https://github.com/messense/jieba-rs) 和 [DCjanus/cang-jie](https://github.com/DCjanus/cang-jie) 作为分词.

## 整体感受

代码并不复杂, 简单的启动一个 web 容器, 加载 csv 构建索引,
提供一个查询接口.

代码很整洁, 看着很舒服, 很多语法并不了解, 但可以猜测大概是什么意思, 并不影响阅读.

项目是对 zlib 的内容检索, 所以最重要作者整理的那几个 csv 的文件.
既然是 csv 文件, 对我来说不需要部署这个项目了, 直接 grep 就搞定了, 就是慢一点而已.
