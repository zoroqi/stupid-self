package q500

type nge struct {
	i int
	v int
}

func NextGreaterElements2(nums []int) []int {
	r := make([]int, len(nums))
	var stack []nge
	push := func(n nge) {
		stack = append(stack, n)
	}
	pop := func() nge {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return s
	}
	peek := func() nge {
		return stack[len(stack)-1]
	}
	isEmpty := func() bool {
		return len(stack) == 0
	}
	l := len(nums)
	for i := 0; i < l*2; i++ {
		num := nums[i%l]
		if isEmpty() {
			if i < l {
				push(nge{i: i, v: num})
			}
		} else {
			for !isEmpty() {
				n := peek()
				if n.v >= num {
					break
				} else {
					r[n.i] = num
					pop()
				}
			}
			if i < l {
				push(nge{i: i, v: num})
			}
		}
	}
	for !isEmpty() {
		p := pop()
		r[p.i] = -1
	}
	return r
}

func NextGreaterElementsSimple(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return nums
	}
	r := make([]int, len(nums))
Outer:
	for i, n := range nums {
		for j := 0; j < l; j++ {
			n2 := nums[(i+j)%l]
			if n < n2 {
				r[i] = n2
				continue Outer
			}
		}
		r[i] = -1
	}
	return r
}
