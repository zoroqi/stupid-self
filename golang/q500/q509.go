package q500

// 仿照haskell写的
func Fib_haskell(N int) int {
	fibs := []int{1, 1}
	fibsTemp := fibs
	take := func() {
		fibs = append(fibs, fibsTemp[0]+fibsTemp[1])
		fibsTemp = fibs[len(fibs)-2:]
	}
	for i := 0; i < N-2; i++ {
		take()
	}
	return fibs[len(fibs)-1]
}

func Fib(N int) int {
	if N == 0 {
		return 0
	}
	if N <= 2 {
		return 1
	}
	f1 := 1
	f2 := 1
	for i := 0; i < N-2; i += 2 {
		f1 = f1 + f2
		f2 = f1 + f2
	}
	if N%2 == 1 {
		return f1
	}
	return f2
}

func Fib_recursion(N int) int {
	var f func(int) int
	f = func(i int) int {
		if i == 0 {
			return 0
		}
		if i <= 2 {
			return 1
		}

		return f(i-1) + f(i-2)
	}
	return f(N)
}
