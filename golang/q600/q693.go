package q600

/**
# 693. 交替位二进制数
```
给定一个正整数，检查他是否为交替位二进制数：换句话说，就是他的二进制数相邻的两个位数永不相等。

示例 1:

输入: 5
输出: True
解释:
5的二进制数是: 101
示例 2:

输入: 7
输出: False
解释:
7的二进制数是: 111
示例 3:

输入: 11
输出: False
解释:
11的二进制数是: 1011
 示例 4:

输入: 10
输出: True
解释:
10的二进制数是: 1010

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-number-with-alternating-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

## 愚笨的思考

第一次想到逐个向右移动并累加计算. 转念这种有规律的二进制数应该通过位运算直接出结果

右移以为后异或整个数字变为全1, +1在做与运算结果为0. 而不是交替出现的结果就不是0了
```
n    001010
n>>1 000101
^    111111
+1 ^ 000000
```

 */
func HasAlternatingBits(n int) bool {
	return (n^(n>>1))&(n^(n>>1)+1) == 0
}