---
id: 20221012221549_8bb8bdbb3b374390
---

# decimal

## 代码统计

[github](https://github.com/shopspring/decimal)

一个golang 高性能、任意精度、浮点十进制库.

从实现上看是可以保障线程安全的, 所有数据都是不可变的.

```
      22 text files.
      22 unique files.
       5 files ignored.

github.com/AlDanial/cloc v 1.84  T=0.07 s (270.4 files/s, 80410.1 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6            508            676           3786
XML                              9              0              0            220
Markdown                         2             46              0            103
YAML                             1              3              0             10
-------------------------------------------------------------------------------
SUM:                            18            557            676           4119
-------------------------------------------------------------------------------
```

复杂度

```
6 file analyzed.
==============================================================
NLOC    Avg.NLOC  AvgCCN  Avg.token  function_cnt    file
--------------------------------------------------------------
    944      10.4     2.6       78.5        85     ./decimal.go
     75      67.0    28.0      419.0         1     ./rounding.go
    163       9.1     2.5       64.1        17     ./decimal_bench_test.go
    311      18.0     5.2      107.3        13     ./decimal-go.go
   2271      27.5     5.4      189.0        41     ./decimal_test.go
```


## 核心结构

```golang
// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	value *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	exp int32 // 科学记数法中的指数部分
}
```

更核心的结构是 `big.Int` golang原生的包, 这里为啥不直接用`big.Float`作为支撑呢?   看完基本代码后的猜测原因:

1. 精度太过复杂
2. 性能不好
3. 表示复杂, 不易于扩展
4. 中间会涉及到IEEE中特殊值状态判断, 比如NaN, Inf这种

这里的结构和[[IEEE-754]]是一致的, 不过采用的10进制存储, [[IEEE-754]]是二进制.  

## 数据构造

可以通过字符串, int, float进行构造. 

int比较简单, 直接big.NewInt就可以了.

string对科学计数法做了特殊处理,  对e做了特殊处理. 科学记数法保证e的指数部分必须是整数.  针对实数部分取掉小数点进行组合就

float处理比较复杂

这里的复杂计算在于对`float64`中精度进行计算[[IEEE-754]]. 简单来说没看懂, 之后有时间继续


## 数学计算

基于以上结构, 我个人的实现逻辑. 

big.Int提供了基础所有操作, 只要对exp和num作对正确的运算就好.  构造的时候将所有数字转换成了整型. 剩下要处理的就是指数部分

### 我的计算

#### 加减
1. 将指数部分变为相同的
2. 进行加减法
3. 指数部分选择小的那个

```
num1, num2

sub_exp = |num1.exp-num2.exp| // 两个位数的差值

以num1 > num2 为例子

newNum.value = num1.value*10^num1.exp +/- num2.value

newNum.exp = min(num1.exp, num2.exp)

```

#### 乘
1. 数字部分直接相乘
2. 指数部分相加就好

#### 除

不会, 这里涉及到了精度问题. (即使不考虑精度我也不会)

### 实际算法

#### 加减

1. 进行数据缩放
    1. 选取小的指数作为新的指数部分
    2. 对较大的数字进行变基
    3. 变基
        1. 两个指数相减的绝对值
        2. 对指数做相应调整.
2. 执行加减法

这个和我设想的是一致的, 没有太大区别

#### 乘法

和我的逻辑一致, 多做了一个越界判断, exp是`int32`

#### 除法

默认精度是16位

1. 指数部分计算num1.exp-num2-exp+16
2. 计算余数部分的指数值
3. 使用big里的QueRem进行计算
4. 针对余数进行操作, 生成新的余数
    1. 新余数=余数*2
5. 新余数大于除数进行精度校准
    1. 除数和被除数符号相反,  `商-1*10^-16`
    2. 除数和被除数符号相通或一个等于0,`商+1*10^-16`

核心逻辑
1. 现将数字放到足够大
2. 进行除法运算, 获得商和余数
3. 步骤4\~5 主要作用是对精度后的最后一位进行四舍五入
    * 这里第一次看的时候没很理解在干啥. 

四舍五入的原理, 以除数20为例, 余数只能是1\~19. 需要进位的是10\~19的部分. 既是$$r\div20\geq0.5$$的部分, 这个里可以等价为$$(r*2)\geq20$$


### 剩下的计算方法

sin, cos, tan, atan就不看了, 我是一点都不会. 估计会用到某种**复变函数或积分变换**的只是吧


## 几个golang的简单比较

[ericlagergren/decimal](https://github.com/ericlagergren/decimal)

```
     107 text files.
     106 unique files.
       9 files ignored.

github.com/AlDanial/cloc v 1.84  T=0.42 s (239.4 files/s, 60399.4 lines/s)
--------------------------------------------------------------------------------
Language                      files          blank        comment           code
--------------------------------------------------------------------------------
Go                               89           1450           2274          21001
Python                            2             50             14            336
Markdown                          3             32              0            111
Java                              1             12              1             86
C                                 1              4              0             40
YAML                              1              8              0             29
Bourne Again Shell                3              6              0             18
XML                               1              0              0             11
--------------------------------------------------------------------------------
SUM:                            101           1562           2289          21632
--------------------------------------------------------------------------------
```

直接涉及到2w行的代码, 不是很适合直接使用, 依赖太多了. 

[apd](https://github.com/cockroachdb/apd)

```
      51 text files.
      51 unique files.
      29 files ignored.

github.com/AlDanial/cloc v 1.84  T=0.07 s (325.8 files/s, 82993.4 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              21            516           1047           4258
Markdown                         1              9              0             15
YAML                             1              3              0             11
-------------------------------------------------------------------------------
SUM:                            23            528           1047           4284
-------------------------------------------------------------------------------
```


这个项目很好玩, 将数据拆成了两部分, `Decimal`没有二元运算(加减乘除), 只有一元运算(绝对值等)

使用了一个Context进行相关的计算. 很迷的依旧是一种入参当出参的代码设计. 

细看逻辑是类似的, 但实现不是我喜欢的风更. 

decimal的实现所有计算中的错误采用painc进行抛出, 而apd采用error方案, 这里我赞成error方案. painc主要是产生原因是超出exp(int32)的范围导致的, 而这种超出范围应作为error更为合适, 让开发主动处理错误
