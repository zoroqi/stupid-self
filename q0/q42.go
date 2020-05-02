package q0

func Trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	maxHeightIndex := 0
	max := height[0]
	for i, h := range height {
		if h > max {
			max = h
			maxHeightIndex = i
		}
	}
	sum := 0
	before := 0
	for i := 0; i <= maxHeightIndex; i++ {
		if height[before] <= height[i] {
			beforeHeight := height[before]
			for j := before; j < i; j++ {
				sum += beforeHeight - height[j]
			}
			before = i
		}
	}
	before = len(height) - 1
	for i := before; i >= maxHeightIndex; i-- {
		if height[before] <= height[i] {
			beforeHeight := height[before]
			for j := before; j > i; j-- {
				sum += beforeHeight - height[j]
			}
			before = i
		}
	}
	return sum
}

func FirstTrap(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	stack := []int{0}

	offset := 0

	push := func(num int) {
		offset += 1
		if len(stack) > offset {
			stack[offset] = num
		} else {
			stack = append(stack, num)
		}
	}

	pop := func() int {
		offset -= 1
		return stack[offset+1]
	}

	getTop := func() int {
		return stack[offset]
	}

	sum := 0
	for i, h := range height {
		if height[getTop()] < h {
			before := pop()
			beforeHeight := height[before]
			for j := before; j < i; j++ {
				sum += beforeHeight - height[j]
			}
			push(i)
		}
	}

	stack[0] = len(height) - 1
	offset = 0

	for i := len(height) - 1; i >= 0; i-- {
		h := height[i]
		if height[getTop()] <= h {
			before := pop()
			beforeHeight := height[before]
			for j := before; j > i; j-- {
				sum += beforeHeight - height[j]
			}
			push(i)
		}
	}

	return sum
}
