package q400

func nextGreaterElementPlanB(nums1, nums2 []int) []int {
	if len(nums2) == 0 || len(nums1) == 0 {
		return []int{}
	}

	queue := make([]int, 0)
	empty := func() bool {
		return len(queue) == 0
	}
	enqueue := func(num int) {
		queue = append(queue, num)
	}
	dequeue := func() {
		if !empty() {
			queue = queue[1:]
		}
	}
	frontqueue := func() int {
		return queue[0]
	}

	enqueue(nums2[0])
	newNums2 := make(map[int]int, len(nums2))
	for _, n2 := range nums2[1:] {
		if !empty() {
			n := frontqueue()
			for n < n2 {
				newNums2[n] = n2
				dequeue()
				if empty() {
					break
				}
				n = frontqueue()
			}
		}
		enqueue(n2)
	}
	for !empty() {
		newNums2[frontqueue()] = -1
		dequeue()
	}
	result := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		result = append(result, newNums2[v])
	}
	return result
}

func nextGreaterElementPlanC(nums1, nums2 []int) []int {
	if len(nums2) == 0 || len(nums1) == 0 {
		return []int{}
	}
	stack := make([]int, 0)
	empty := func() bool {
		return len(stack) == 0
	}
	pop := func() int {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return n
	}
	peek := func() int {
		return stack[len(stack)-1]
	}
	push := func(n int) {
		stack = append(stack, n)
	}

	newNums2 := make(map[int]int, len(nums2))
	for _, n2 := range nums2 {
		if !empty() {
			n := peek()
			for n < n2 {
				newNums2[n] = n2
				pop()
				if empty() {
					break
				}
				n = peek()
			}
		}
		push(n2)
	}
	for !empty() {
		newNums2[peek()] = -1
		pop()
	}
	result := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		result = append(result, newNums2[v])
	}
	return result
}