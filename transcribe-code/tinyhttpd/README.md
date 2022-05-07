# Tintyhttpd

源码地址: [EZLippi/Tinyhttpd: Tinyhttpd](https://github.com/EZLippi/Tinyhttpd)

源码下载: [tiny-httpd - Browse /tiny-httpd at SourceForge.net](https://sourceforge.net/projects/tiny-httpd/files/tiny-httpd/)

源码很少, 虽然是 C 的, 但看懂问题应该不大.

## 代码基本情况

### 行数统计

代码基本情况, 统计工具[coca](https://github.com/inherd/coca)
```

───────────────────────────────────────────────────────────────────────────────
Language                 Files     Lines   Blanks  Comments     Code Complexity
───────────────────────────────────────────────────────────────────────────────
C                            2       549       48        91      410         85
HTML                         1        11        0         0       11          0
License                      1       674      121         0      553          0
Makefile                     1         9        1         0        8          0
Markdown                     1       101       14         0       87          0
───────────────────────────────────────────────────────────────────────────────
Total                        6      1344      184        91     1069         85
───────────────────────────────────────────────────────────────────────────────
Estimated Cost to Develop $28,974
Estimated Schedule Effort 3.580819 months
Estimated People Required 0.718878
───────────────────────────────────────────────────────────────────────────────
Processed 59346 bytes, 0.059 megabytes (SI)
───────────────────────────────────────────────────────────────────────────────
```

### call graph

方法调用图, 使用工具 [cally](https://github.com/chaudron/cally), 工具依赖 gcc和graphviz

完整调用图
![](./full_call_graph.svg)

内部函数调用图
![](./inner_call_graph.svg)


## 辅助网站

* [The Linux man-pages project](https://www.kernel.org/doc/man-pages/) 库帮助, [Linux manual pages: section 3](https://man7.org/linux/man-pages/dir_section_3.html)
* [RFC 1945: Hypertext Transfer Protocol -- HTTP/1.0](https://www.rfc-editor.org/rfc/rfc1945.html)

## 无法理解和个人解答

作为一个没有经过 C/C++ 洗礼的程序员, 无法理解其中的书写方式和设计方案.

1. C 语言为啥用通过修改入参实现对象的修改
    1. 大概和 C 语言本身没有错误处理机制
    2. 内存分配有关
2. 为啥要的定义 `char[1024]` 但不写满\(尽可能满一些, 或者至少写一半\)就调用 send 方法.
    1. 主要针对的是 400, 404, 500 这几段处理, 写到`\r\n` 就发送了.
    2. 猜测原因
        1. 一种习惯
        2. RFC 文档中有某种隐含规定, 但我没看出来
3. get_line 中的 recv 调用一次只读取一个字符是不是有点少了.
    1. 后边想了一下, 可能真的需要一个一读取
    2. 因为是流, 消费超过`\r\n`数据没有办法二次消费.
        * 只能自己进行缓存
        * 这样写代码复杂度会复杂很多
    3. 每一次消费都使用 MSG_PEEK 效果还没有一个一个来的好.
4. headers 方法要把 filename 传进去, 但是为什么呢?
    1. 感觉是为了扩展留下的
5. `const char *`是可以指代数组吗?
    1. 可以转换, 但不别扭吗. https://www.cnblogs.com/kevinWu7/p/10163446.html
    2. 这种细小的差异真的是蛋疼啊
6. 在执行 CGI 代码后如何把执行结果发送出去?
    1. 看代码是通过命令行执行 CGI
    2. 使用 pipe 进行数据传输
    3. 通过代码逻辑上 if/else 中执行 CGI 和发送结果是在两个分支中.
    4. 这里的串联使用 fork 来实现.
        1. fork 可以返回两次, 一次是执行 true 分支, 一次执行 false 分支
        2. fork 被调用一次，返回两次，一次在父进程中返回子进程 PID ，一次在子进程中返回0。 fork 失败返回负数，发生在    PID 个数达上限或内存不足时。
        3. fork 复制时复制了父进程的堆栈段, 所以执行逻辑是从 fork 函数执行的地方开始分叉的. 也就是 if / else 逻辑是是在两个 process 执行的.

### 新知

fork 是个神奇的东西, 没写过 C/C++ 和关于 Linux 的代码, 完全没想到还有这种代码执行方式. 我个人开发和使用的语言是 java 和 golang 是无法从一个方法中间开始执行代码的. 但 fork 的复制机制实现了, 子 process 的执行开始点是方法的中间. 这让我想到了 goto 这个无条件跳转, 但 fork 是吧整个栈复制一份后再跳转.

## 代码流程

1. 监听套接字
2. 收到监听创建线程进行处理
3. 使用 CGI 进行处理.
4. 结束
