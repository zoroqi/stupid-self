---
aliases:
- 1912. 设计电影租借系统
- 1912. design movie rental system
tc:
- leetcode
- algorithm
leetcode:
  num: 1912
  url: https://leetcode.cn/problems/design-movie-rental-system/
date: "2023-05-23"
id: 20230523082540_5719d51c93314139
---

# 1912. 设计电影租借系统

你有一个电影租借公司和 n 个电影商店。你想要实现一个电影租借系统，它支持查询、预订和返还电影的操作。同时系统还能生成一份当前被借出电影的报告。

所有电影用二维整数数组 entries 表示，其中 `entries[i] = [shopi, moviei, pricei]` 表示商店 shopi 有一份电影 moviei 的拷贝，租借价格为 pricei 。每个商店有**至多一份**编号为 moviei 的电影拷贝。

系统需要支持以下操作：

- Search：找到拥有指定电影且 **未借出** 的商店中 最便宜的 5 个 。商店需要按照 价格 升序排序，如果价格相同，则 shopi 较小 的商店排在前面。如果查询结果少于 5 个商店，则将它们全部返回。如果查询结果没有任何商店，则返回空列表。
- Rent：从指定商店借出指定电影，题目保证指定电影在指定商店未借出 。
- Drop：在指定商店返还 之前已借出 的指定电影。
- Report：返回 最便宜的 5 部已借出电影 （可能有重复的电影 ID），将结果用二维列表 res 返回，其中 `res[j] = [shopj, moviej]` 表示第 j 便宜的已借出电影是从商店 shopj 借出的电影 moviej 。res 中的电影需要按 价格 升序排序；如果价格相同，则 shopj 较小 的排在前面；如果仍然相同，则 moviej 较小 的排在前面。如果当前借出的电影小于 5 部，则将它们全部返回。如果当前没有借出电影，则返回一个空的列表。
请你实现 MovieRentingSystem 类：

- `MovieRentingSystem(int n, int[][] entries)` 将 MovieRentingSystem 对象用 n 个商店和 entries 表示的电影列表初始化。
- `List<Integer> search(int movie)` 如上所述，返回 未借出 指定 movie 的商店列表。
- `void rent(int shop, int movie)` 从指定商店 shop 借出指定电影 movie 。
- `void drop(int shop, int movie)` 在指定商店 shop 返还之前借出的电影 movie 。
- `List<List<Integer>> report()` 如上所述，返回最便宜的 已借出 电影列表。

注意：测试数据保证 rent 操作中指定商店拥有 未借出 的指定电影，且 drop 操作指定的商店 之前已借出 指定电影。


```
示例 1：
示例 1：

输入：
["MovieRentingSystem", "search", "rent", "rent", "report", "drop", "search"]
[[3, [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]]], [1], [0, 1], [1, 2], [], [1, 2], [2]]
输出：
[null, [1, 0, 2], null, null, [[0, 1], [1, 2]], null, [0, 1]]

解释：
MovieRentingSystem movieRentingSystem = new MovieRentingSystem(3, [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]]);
movieRentingSystem.search(1);  // 返回 [1, 0, 2] ，商店 1，0 和 2 有未借出的 ID 为 1 的电影。商店 1 最便宜，商店 0 和 2 价格相同，所以按商店编号排序。
movieRentingSystem.rent(0, 1); // 从商店 0 借出电影 1 。现在商店 0 未借出电影编号为 [2,3] 。
movieRentingSystem.rent(1, 2); // 从商店 1 借出电影 2 。现在商店 1 未借出的电影编号为 [1] 。
movieRentingSystem.report();   // 返回 [[0, 1], [1, 2]] 。商店 0 借出的电影 1 最便宜，然后是商店 1 借出的电影 2 。
movieRentingSystem.drop(1, 2); // 在商店 1 返还电影 2 。现在商店 1 未借出的电影编号为 [1,2] 。
movieRentingSystem.search(2);  // 返回 [0, 1] 。商店 0 和 1 有未借出的 ID 为 2 的电影。商店 0 最便宜，然后是商店 1 。
```

**提示：**

提示：

- `1 <= n <= 3 * 10^5`
- `1 <= entries.length <= 10^5`
- `0 <= shopi < n`
- `1 <= moviei, pricei <= 10^4`
- 每个商店 至多 有一份电影 moviei 的拷贝。
- search，rent，drop 和 report 的调用总共不超过 105 次。

## 愚笨的思考

感觉是一个总和性质的题, 主要是如何创建一个合理的索引, 十分想尝试是不是可以引入 sqlite.
如果可以引入 sqlite 可能解决起来会很简单.
测试发现不能引入 `github.com/mattn/go-sqlite3`.

### 直接实现 PlanA

只能自己实现了, 这个代码用 haskell 是没戏了, 主要是不知道如何实现交换设计和状态的保持(State Monad 还没有看懂).

简单看需求可以真的, sql 都写完了, 需要代码实现这些, 真的有点蛋疼, 还是 sql 方便啊.

- `select * from movie_renting_system where movie = ? and rented = 0 order by price, shop limit 5;`
- 借出/还入 `update movie_renting_system set rented = (1/0) where shop = ? and movie = ?;`
- `select * from movie_renting_system where rented = 1 order by price, shop, movie limit 5;`

字段除了 id 包含四个字段
- shop
- movie
- rented 借出状态
- price 价钱

需要考虑的索引就是 rented, movie, 整体排序采用 shop 进行排序可能是一个不错的选择. 还是想用 sql ...

就是各种 map 嵌套, 要不就是循环, 反正也没有什么好办法, 以嵌套为主吧.

然后失败, 在第 41 个测试用例失败了, 问题是不告诉我是那个具体结果是什么. 我自己连测试用例都没有办法写...

计算错误计数方式, 导致的问题.

实现数据库的人真的是伟大啊, 完美解决这种问题.
