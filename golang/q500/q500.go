package q500

/**
# 500
给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。


![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/keyboard.png)


示例：

输入: ["Hello", "Alaska", "Dad", "Peace"]
输出: ["Alaska", "Dad"]


注意：

你可以重复使用键盘上同一字符。
你可以假设输入的字符串将只包含字母。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/keyboard-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 愚笨的思考

建立一个映射表, 查询的对应映射的行数, 同一行返回true
*/

var key_mapping = []int{
	1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2,
	//a,b,c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z
}

func FindWords(words []string) []string {
	var r []string

	findLine := func(c uint8) int {
		if c >= 97 {
			c = c - 97
		} else {
			c = c - 65
		}
		return key_mapping[c]
	}
Outer:
	for _, w := range words {
		lineNum := findLine(w[0])
		for _, b := range w {
			if findLine(uint8(b)) != lineNum {
				continue Outer
			}
		}
		r = append(r, w)
	}
	return r
}
