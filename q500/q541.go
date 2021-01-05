package q500

/**
# 541. 反转字符串 II

给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。


示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"


提示：

该字符串只包含小写英文字母。
给定字符串的长度和 k 在 [1, 10000] 范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 愚笨的思考

没啥逻辑, 就是处理好分支逻辑就好了.
*/

func ReverseStr(s string, k int) string {
	if k == 1 {
		return s
	}

	reverse := func(bb []byte) {
		l := len(bb) - 1
		i := 0
		// 逻辑更简单
		for i < l {
			bb[i], bb[l] = bb[l], bb[i]
			l--
			i++
		}
	}
	//reverse := func(bb []byte) {
	//	ll := len(bb) / 2
	//	l := len(bb) - 1
	//	for i := 0; i < ll; i++ {
	//		bb[i], bb[l-i] = bb[l-i], bb[i]
	//	}
	//}
	bs := []byte(s)
	l := len(bs)
	for i := 0; i < l; i = i + 2*k {
		max := i + k
		if max >= l {
			max = l
		}
		reverse(bs[i:max])
	}

	return string(bs)
}
