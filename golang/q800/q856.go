package q800

import (
	. "github.com/zoroqi/stupid-self/golang"
)

func scoreOfParenthesesPlanA(s string) int {
	nums := Stack[int]{}
	for _, c := range s {
		if c == '(' {
			nums.Push(0)
		} else {
			sum := 0
			top, _ := nums.Pop()
			if top == 0 {
				sum += 1
			} else {
				sum = top * 2
				next, _ := nums.Pop()
				if next != 0 {
					sum += next
				}
			}
			for !nums.IsEmpty() {
				n, _ := nums.Pop()
				if n == 0 {
					nums.Push(0)
					break
				} else {
					sum += n
				}
			}
			nums.Push(sum)
		}
	}
	n1, _ := nums.Pop()
	return n1
}
