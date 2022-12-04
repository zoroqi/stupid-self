---
id: 20221012195333_2c465e25d1fe4693
date: "2022-10-12"
github: https://github.com/SmileXie/cdls
aliases:
- cdls 源码阅读
tags:
- 源码
- rust
tc:
- github
- reading_code
---

# cdls 源码阅读

[cdls](https://github.com/SmileXie/cdls)

## 简介

一个简单的 rust 库, 一个小命令行工具, 用来代替 cd 和 ls 两个命令.
工具对我来时没有多大意义,
    我习惯了 z, cd, tree, ll, grep 的混合操作,
    对这部分交互性质的 cli 命令兴趣不大.

没有学过 rust 代码, 所以对 rust 好奇, 这个项目刚好可以看看. 感受下.

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Cargo Lock                   1       391       45         2      344          0
License                      1       674      121         0      553          0
Markdown                     1        68       28         0       40          0
Rust                         1       807      118        13      676        128
TOML                         1        20        2         3       15          0
gitignore                    1         2        0         0        2          0
───────────────────────────────────────────────────────────────────────────────
Total                        6      1962      314        18     1630        128
───────────────────────────────────────────────────────────────────────────────
```

## 源码

rust 完全没有学过, 但看懂代码不是很费劲.
根据我从知乎等平台上的了解,
    rust 的语法复杂在设计上,
    对阅读影响不大.
比如: 借用, 引用, 转移完全不懂,
    但不影响看,
    当做都是赋值好了, 反正就是赋值上的一些限制.

代码中对我最大的一个新发现是,
    rust 可以对基础类型或其它库的类型做接口扩展.

```rust
trait I32Ext {
    fn to_char(&self) -> char;
}
impl I32Ext for i32 {
    fn to_char(&self) -> char {
        let tmp: u8;
        if self < &0 {
            tmp = 0;
        } else ...
        return tmp as char;
}
```

我会 java 和 golang 这个都是不允许的.
java 想扩展一些内容需要通过继承来实现,
    而基础类型和包装类都不能继承.
golang 语法好像可以, 但有限制不允许对别的包进行扩展.
可以对类型创建一个新的别名在实现也是可以的,
    但这和 rust 还是有些不同.    
但我没有在 rust 简单教程找到相关限制约束条件,
    不知道这种扩展可以影响范围是什么,
        是只有本包才会有作用,
        还是导入本包的都会有作用.
印象中 haskell 是可以这么扩展的.

这种扩展很有意思, 感觉可以用这种特性玩出很多花来.
我第一反应就是这种扩展方式可以在项目中引入一些特别的复杂性,
    某个基础类型在处理中不知道为啥多出点特别功能,
    语法不熟的就可以找一段时间.
我猜测这种影响范围应该是只限制在包内, 无法传播的,
    不然就像一个搞破坏的 ts 库一样了, 可以直接把整个项目搞崩.
但我依旧感觉这个特性很好.

## 对这个特性的使用

```rust
trait PathBufExt {
    fn file_size(&self) -> u64;
    fn file_modified_time(&self) -> DateTime<Utc>;
    fn file_type(&self) -> &str;
    fn fuzzy_search_score(&self, search_str: &str) -> f32;
}

impl PathBufExt for PathBuf {
    fn file_size(&self) -> u64 { ... }
    fn file_modified_time(&self) -> DateTime<Utc> { ... }
    fn file_type(&self) -> &str { ... }
    fn fuzzy_search_score(&self, search_str: &str) -> f32 { .... }
}
```

通过这种方式来扩展 PathBuf,
    扩展出来的方法主要用于了展示中的排序.

我个人实现的话会用一组函数来处理, 并不会去选择扩展 PathBuf.
```rust
fn file_size(path: &PathBuf) -> u64 { ... }
...
```
当然这两种方式在这种规模下不会产生差别,
    但项目在大几倍作者这种扩展方式可能会更好.
没有增加新的类型, 也有更好的内聚性.