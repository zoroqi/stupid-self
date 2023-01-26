package q500

/*
*
# 551 学生出勤记录 I

给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符：

1.  **'A'** : Absent，缺勤
2.  **'L'** : Late，迟到
3.  **'P'** : Present，到场

如果一个学生的出勤记录中不**超过一个'A'(缺勤)**并且**不超过两个*连续*的'L'(迟到)**,那么这个学生会被奖赏。

你需要根据这个学生的出勤记录判断他是否会被奖赏。

**示例 1:**

**输入:** "PPALLP"
**输出:** True

**示例 2:**

**输入:** "PPALLL"
**输出:** False

## 愚笨的思考

问题可以理解为, 最多出现两次A或连续三次L就返回false

进行简单计数即可, 连续需要通过再非L情况小改为0
*/
func CheckRecord(s string) bool {
	a, l := 0, 0
	for _, v := range s {
		switch v {
		case 'A':
			a++
			l = 0
		case 'L':
			l++
		default:
			l = 0
		}
		if a >= 2 || l >= 3 {
			return false
		}
	}
	return true
}

/*
*
# 557. 反转字符串中的单词 III

给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

**示例：**

**输入：**"Let's take LeetCode contest"
**输出：**"s'teL ekat edoCteeL tsetnoc"

********提示：********

*   在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

## 愚笨的思考

遍历遇到空格翻转之前的数据
*/
func ReverseWords(s string) string {
	newS := []byte(s)
	start := 0
	reverse := func(i, j int) {
		for i < j {
			newS[i], newS[j] = newS[j], newS[i]
			i++
			j--
		}
	}
	for i, v := range newS {
		if v == ' ' {
			reverse(start, i-1)
			start = i + 1
		}
	}
	reverse(start, len(newS)-1)
	return string(newS)
}
