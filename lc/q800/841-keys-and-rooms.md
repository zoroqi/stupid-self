---
aliases:
- 841. 钥匙和房间
- 841. keys and rooms
tc:
- leetcode
- algorithm
leetcode:
    num: 841
    url: https://leetcode.cn/problems/keys-and-rooms/
    tags:
    - dfs
    - bfs
    - 图
    - 有向图
---

# 841. 钥匙和房间

有 `n` 个房间，房间按从 `0` 到 `n - 1` 编号。最初，除 `0` 号房间外的其余所有房间都被锁住。你的目标是进入所有的房间。然而，你不能在没有获得钥匙的时候进入锁住的房间。

当你进入一个房间，你可能会在里面找到一套不同的钥匙，每把钥匙上都有对应的房间号，即表示钥匙可以打开的房间。你可以拿上所有钥匙去解锁其他房间。

给你一个数组 `rooms` 其中 `rooms[i]` 是你进入 `i` 号房间可以获得的钥匙集合。如果能进入 **所有** 房间返回 `true`，否则返回 `false`。

```
示例 1：

输入：rooms = [[1],[2],[3],[]]
输出：true
解释：
我们从 0 号房间开始，拿到钥匙 1。
之后我们去 1 号房间，拿到钥匙 2。
然后我们去 2 号房间，拿到钥匙 3。
最后我们去了 3 号房间。
由于我们能够进入每个房间，我们返回 true。

示例 2：

输入：rooms = [[1,3],[3,0,1],[2],[0]]
输出：false
解释：我们不能进入 2 号房间。
```

**提示：**

* `n == rooms.length`
* `2 <= n <= 1000`
* `0 <= rooms[i].length <= 1000`
* `1 <= sum(rooms[i].length) <= 3000`
* `0 <= rooms[i][j] < n`
* 所有 `rooms[i]` 的值 **互不相同**

## 愚笨的思考

### 暴力方案 PlanA

示例二, 可以理解是有向图的遍历, 判断是否遍历完所有节点. 使用了队列来遍历, 可能栈更好, 但这个没有他打区别. 当然可能递归也很简单.

### 递归遍历 PlanB

golang 的递归比较简单, 但 haskell 如何实现遍历, 我需要一个变量存储遍历过的节点, 问题是这个要怎么做? 和 golang 使用相同的方式用一个数组存储, haskell 修改列表中的元素比较费劲, 需要别的方式, 可以直接用 Set 进行存储, 但感觉不是很好.

剩下就是 haskell 咋写图的遍历了. 防止循环导致无限递归, 必须使用一个 Set 结构进行存储, 最后写出一个复杂的函数来进行存储.

算法基本状态, 使用深度优先遍历, 返回遍历过的节点 Set; 合并自身和所有子调用返回作为返回结果. 但是递归调用让我很不爽, 感觉效果

网上找到的一个方案, 使用了一个我没用过的包 "Data.Sequence", 作为队列但这个没有用过.

```hs
module Main where

import qualified Data.HashMap.Strict as HM
import           Data.Maybe (fromJust)
import qualified Data.Sequence       as DS

graph :: HM.HashMap String [String]
graph = HM.fromList [("you",["alice", "bob", "claire"]),
                    ("bob", ["anuj", "peggy"]),
                    ("alice", ["peggy"]),
                    ("claire", ["thom", "jonny"]),
                    ("anuj",[]),
                    ("peggy",[]),
                    ("thom",[]),
                    ("jonny",[])
                   ]

personIsSeller :: String -> Bool
personIsSeller name = last name == 'm'

search :: HM.HashMap String [String] -> String -> Bool
search graph name = loop $ DS.fromList (graph HM.! name)
  where loop queue
          | null queue = False
          | personIsSeller h = True
          | otherwise = loop $ (DS.drop 1 queue) DS.>< DS.fromList (graph HM.! h)
          where h = queue `DS.index` 0

main :: IO ()
main = do
  print $ search graph "you"
```
