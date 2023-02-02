---
id: 20230202153850_874acd40c2144cef
date: "2023-02-02"
aliases:
- dragon book appendix A
- 龙书附录代码 A
tags:
- 源码
- java
tc:
- reading_code
---

# 龙书附录代码 A

源码可以在这里下载到 [Compilers: Principles, Techniques, and Tools (Dragon Book)](https://suif.stanford.edu/dragonbook/)  或 [lu1s/dragon-book-source-code](https://github.com/lu1s/dragon-book-source-code)

## 简介

代码是针对 [[编译原理]] 中第 2.5 到 2.8 节的代码示例. 功能是对文本内容进行代解析, 并生成 #编译/三地址码 .

```
───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
Java                        34      1071      203        13      855        190
Swig                        30       662        0         0      662          0
XML                          3        22        0         0       22          0
Plain Text                   2        29        0         0       29          0
Makefile                     1        27        4         0       23          0
gitignore                    1         8        0         3        5          0
───────────────────────────────────────────────────────────────────────────────
Total                       71      1819      207        16     1596        190
───────────────────────────────────────────────────────────────────────────────
```

## 简单使用使用

需要安装 jdk, 执行 make compile 进行编译.

### 测试
测试代码 `test.txt`
```
{
    int i; int j; float v; float x; float[100] a;
    while(true) {
        do i = i+1; while(a[i]<v);
        do j = j+1; while(a[j]>v);
        if (i>=j) break;
        x = a[i]; a[i]=a[j];a[j]=x;
    }
}
```

执行 `cat test.txt | java main.Main` 可以执行代码.
输出结果
```
L1:L3:	i = i + 1
L5:	t1 = i * 8
	t2 = a [ t1 ]
	if t2 < v goto L3
L4:	j = j + 1
L7:	t3 = j * 8
	t4 = a [ t3 ]
	if t4 > v goto L4
L6:	iffalse i >= j goto L8
L9:	goto L2
L8:	t5 = i * 8
	x = a [ t5 ]
L10:	t6 = i * 8
	t7 = j * 8
	t8 = a [ t7 ]
	a [ t6 ] = t8
L11:	t9 = j * 8
	a [ t9 ] = x
	goto L1
L2:
```

### 语法规则
```
program -> block
block -> {decls stmts}
decls -> decls decl | ε
decl -> type id;
type -> type [num] | basic
stmts -> stmts stmt | ε

stmt -> loc = bool;
    | if (bool) stmt
    | if (bool) stmt else stmt
    | while (bool) stmt
    | do stmt while (bool);
    | break;
    | block
loc -> loc[bool] | id

bool -> bool || join | join
join -> join && equality | equality
equlity -> equality == rel | equalitye != rel | rel
rel -> expr < expr | expr <= expr | expr >= expr | expr > expr | expr
expr -> expr+term | expr - term | term
term -> term * unary | term/unary | unary
unary -> !unary | -urnary | factor
factor -> (bool) | loc | num | real | true | false
```

## 源码阅读

### 词法规则

这里特别需要注意的是, 变量的声明必须在 block 块中的第一行.
所以这样写会在执行中报错. (特别声明, 代码不支持注释)
```
{
    int i; int j;
    while(true) {
        i = 10 + 20;
        j = i * 2;
        int g;  // 这行会报错
        if (g > 20) {
            break;
        }
    }
}
```

### 代码结构

主要 package 功能
* inter 语法树的节点
* lexer 词法分析
* main 启动函数
* parser 语法解析
* symbols 符号表

### lexer

更能比较简单, 并没有使用正则, 一次向前看一个字符进行循环处理.

### parse

对语法树的实现, 消费 lexer 产生的 token 序列, 生成语法树.

这是一个递归的过程, 但实现完全是按照语法规则来的, 不对照语法规则看, 几乎是看不懂代码在干啥的.

生成的结果就是一个 Stmt 类型为根节点的树

### inter

语法树所有节点的定义, 在 parse 解析完后会构建成一棵完整的树.
通过调用 `gen` 生成三地址码.

代码结构分成 Stmt 和 Expr 两部分, Stmt 表示语句和控制结构, Expr 表示计算和表达式结构.

执行过程依旧是深度优先的遍历打印.

```java
// While 的 gen 很熟
// b: 是 begin 的意思, 这个块的开始
// a: 是 after 的意思, 这个块的结束或下一块的开始
public void gen(int b, int a) {
  after = a;              // save label a
  expr.jumping(0, a);     // 打印 while 的判断逻辑
  int label = newlabel(); // 计算新的跳转地址
  emitlabel(label);       // 输出跳转地址
  stmt.gen(label, b);     // 打印 while 体
  emit("goto L" + b);     // 跳转到开始处, 继续循环
}
```

## 整体的感受

整体代码需要在可以看懂语法规则后进行阅读, 代码没有多复杂.
面向对象的代码风格, 看的时候不要想我在调用的具体是什么, 只要想那就看不懂了.
按照语义去理解, jumping 是出现跳转的时候需要用的, gen 生成之后的指令要用的.

可能作为一个入门的语法解释是一个不错的示例, 复杂的示例可以参考 [jhy/jsoup](https://github.com/jhy/jsoup) 和 [yuin/goldmark](https://github.com/yuin/goldmark). 
我之前还看过 [[goldmark-meta]] 但是 goldmark 就没有看过了.
