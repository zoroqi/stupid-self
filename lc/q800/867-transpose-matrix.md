---
aliases:
- 867. 转置矩阵
- 867. transpose matrix
tc:
- leetcode
- algorithm
leetcode:
  num: 867
  url: https://leetcode.cn/problems/transpose-matrix/
  tags:
  - 数组
  - 矩阵
date: "2022-10-18"
id: 20221018025321_a33411d0a2b6422b
---

# 867. 转置矩阵

给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。

矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
```
示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]
示例 2：

输入：matrix = [[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]
```

提示：

* `m == matrix.length`
* `n == matrix[i].length`
* `1 <= m, n <= 1000`
* `1 <= m * n <= 10^5`
* `-10^9 <= matrix[i][j] <= 10^9`

## 愚笨的思考

golang 很简单, 转置只需要初始化之后交换下标就好了, `i,j -> j,i`

haskell 就比较麻烦了, 没有办法简单赋值, 也不是很好构建数组.
按照 golang 的方案是两个循环嵌套, 到 haskell 就是两个递归嵌套了, 想象都费劲.

想了想发现可以用 zipWith 来减少一个递归, 总比自己写两个递归函数要好很多.
```haskell
transpose (x:xs) = zipWith (:) x (transpose xs)
```

执行步骤
```
[1 2 3]
[4 5 6]
[7 8 9]
一, 读取最后一行(n行), 转换成一个 [[a]]
[7]
[8]
[9]
二, 读取倒数第二行(n-1h行), 转换成一个 [[a]] 并和之前的 [[a]] 进行对应索引的前拼接
[4 7]
[5 8]
[6 9]
三, 反复执行第二步, 知道所有行执行完成
[1 4 7]
[2 5 8]
[3 6 9]
```


睡觉躺床上的时候感觉这个流程是一个 fold 的过程, 尝试看看, 最后成功了.
```haskell
transpose = foldr (zipWith (:)) (repeat [])
```

功能和递归是一样的只是递归由自己实现变为了 foldr 实现.
