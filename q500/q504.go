package q500

import (
	"strconv"
)

/**
# 504. 七进制数

```
给定一个整数，将其转化为7进制，并以字符串形式输出。

示例 1:

输入: 100
输出: "202"
示例 2:

输入: -7
输出: "-10"
注意: 输入范围是 [-1e7, 1e7] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/base-7
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

## 愚笨的思考

进制转换, 10进制转x进制. n%x 获得对应位的数字, n/x减小数据

n位10进制数, 10^(n-1), 10^(n-2), 10^(n-3)...
n位x进制数,  x^(n-1), x^(n-2), x^(n-3)...


*/
func ConvertToBase7(num int) string {
	s := ""
	n := num
	if num < 0 {
		n = -n
	}
	for n != 0 {
		s = strconv.Itoa(n%7)+s
		n = n/7
	}
	if s == "" {
		return "0"
	}
	if num < 0 {
		return "-"+s
	}
	return s
}
