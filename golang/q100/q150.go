package q100

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	var nums []int
	for _, t := range tokens {
		l := len(nums)
		switch t {
		case "+":
			nums = append(nums[:l-2], nums[l-2]+nums[l-1])
		case "-":
			nums = append(nums[:l-2], nums[l-2]-nums[l-1])
		case "*":
			nums = append(nums[:l-2], nums[l-2]*nums[l-1])
		case "/":
			nums = append(nums[:l-2], nums[l-2]/nums[l-1])
		default:
			n, _ := strconv.Atoi(t)
			nums = append(nums, n)
		}
	}
	return nums[0]
}

func first150(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	var stack []int
	push := func(n int) {
		stack = append(stack, n)
	}

	pop := func() int {
		r := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return r
	}

	add := func() int {
		p1 := pop()
		p2 := pop()
		return p1 + p2
	}

	sub := func() int {
		p1 := pop()
		p2 := pop()
		return p2 - p1
	}

	multi := func() int {
		p1 := pop()
		p2 := pop()
		return p1 * p2
	}

	division := func() int {
		p1 := pop()
		p2 := pop()
		return p2 / p1
	}

	for _, t := range tokens {
		switch t {
		case "+":
			push(add())
		case "-":
			push(sub())
		case "*":
			push(multi())
		case "/":
			push(division())
		default:
			n, _ := strconv.Atoi(t)
			push(n)
		}
	}
	return stack[0]
}
