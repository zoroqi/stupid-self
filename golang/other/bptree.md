---
id: 20240307161148_34735b0db76049d4
date: "2024-03-07"
aliases:
- bptree 实现过程中的一些感想
category:
- 算法
tags:
- 算法/B树
tc:
- 算法
---

闲着无聊实现了想实现 B+ 树看看, 发现还是很难的.
即使现在我也不知道自己的逻辑是否是完全正确的, 虽然有单元测试, 但是测试感觉是不够的.
因为是通过随机方案进行的测试, 所以很有可能是错误的(只是跑测试的次数多一点, 但无法确定是否正确)

主要的参考:
- [算法】B树、B+树详解 - H__D - 博客园](https://www.cnblogs.com/h--d/p/14022357.html)
- [B+ 树 - OI Wiki](https://oi-wiki.org/ds/bplus-tree/)
- [周刊（第23期）：图解Blink-Tree：B+Tree的一种并发优化结构和算法 - codedump的网络日志](https://www.codedump.info/post/20220807-weekly-23/)
- [B+树详解 | Ivanzz](https://ivanzz1001.github.io/records/post/data-structure/2018/06/16/ds-bplustree)
- [B+树详解+代码实现（插入篇） - 周小伦 - 博客园](https://www.cnblogs.com/JayL-zxl/p/14304178.html)
- [B+树原理以及Go语言实现 - 个人文章 - SegmentFault 思否](https://segmentfault.com/a/1190000041696709#item-3)
- [B树和B+树的插入、删除图文详解 - nullzx - 博客园](https://www.cnblogs.com/nullzx/p/8729425.html)
- [Introduction of B+ Tree - GeeksforGeeks](https://www.geeksforgeeks.org/introduction-of-b-tree/)

是否是正确的实现? 不是, 我没有实现链表逻辑, 我也不想再补充这部分逻辑了.

当前实现是否还存在问题? 存在一个巨大的问题, 每一个右子树最左侧很有可能会出现一个空节点.
我现在的实现不知道如何规避这个问题, 并且如果将这个空干掉很有可能会出现错误.
我选择使用 `[a,b)` 的索引方案, 这导致在合并和分列的时候, 索引节点中的最左侧节点很可能是空, 这样我就可以实现任何处理只需要在父子节点两层进行交互.
现在暂时不知道如何修正这个问题, 可能我修正需要重新实现整个逻辑, 包括写入逻辑.

在工作中很少实现某个特殊算法, 大部分都是直接调用库, 直接实现发现还是有很多难度的.
上一次在工作中实现的稍微复杂的算法是字典树, 用来实现域名过滤.
在实现这种复杂逻辑的时候, 还是必须补充充足的测试用例, 不然真的不知道自己实现的是否正确.

参考的所有算法讲解, 逻辑上是没有大问题的, 但具体实现上, 问题一大堆.
更重要的是理解的偏差, 导致最初的实现实际上和算法的一些描述有一些偏差, 到最后算法讲解已经没有参考价值了.

这里还有一个空间优化问题, 在一个 m 设置为 10, 按照升序插入 1~1000, 那叶子节点的空间使用只有 50%, 也就是叶子节点只有 5 个元素, 中间节点的使用率也是不足的, 也只有 50%.
那在 mysql 的 B+ 树实现每个页 16kb, 简单的自增插入操作空间使用率就只有 50% 这个太浪费了.
查了资料好像 mysql 针对这个场景有特殊的优化, "[数据库内核月报](http://mysql.taobao.org/monthly/2021/06/05/) [archive](https://web.archive.org/web/20220626080734/http://mysql.taobao.org/monthly/2021/06/05/)" 和 "[从MySQL Bug#67718浅谈B+树索引的分裂优化](https://www.cnblogs.com/mscm/p/13493129.html) [archive](https://web.archive.org/web/20240307092434/https://www.cnblogs.com/mscm/p/13493129.html)", (极端情况下还存在 bug, 蛋疼).
只要没有随机插入, 那可以尝试修改分裂策略; 但是在无法预估场景的情况下, 那 50% 应该是最好的策略; 当然"智能"化的分割可能会更好, 预估数据插入的分布方式后针对性的优化可能会有效果.
