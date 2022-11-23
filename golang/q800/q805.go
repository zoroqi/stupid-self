package q800

func splitArraySameAveragePlanA(nums []int) bool {
	per := permute(nums)
	l := len(nums)
	n, next := per()
	for next {
		for i := 1; i < l; i++ {
			left := avg(n[:i])
			right := avg(n[i:])
			if left == right {
				return true
			}
		}
		n, next = per()
	}
	return false
}

func splitArraySameAveragePlanB(nums []int) bool {
	s := sum(nums)
	l := len(nums)
	average := float64(s) / float64(l)
	pre := splitTwoPlanB(nums)
	n, next := pre(false)
	for next {
		ls := sum(n)
		left := float64(ls) / float64(len(n))
		if left == average {
			n, next = pre(true)
			return true
		}
		n, next = pre(false)
	}
	return false
}

func splitTwoPlanB(nums []int) func(stop bool) ([]int, bool) {
	ch := make(chan []int)
	var dfs func(c int, n []int, temp []int)
	end := false
	dfs = func(c int, n []int, temp []int) {
		if end {
			return
		}
		if len(temp) == c {
			ch <- append([]int{}, temp...)
			return
		}
		for i := 0; i < len(n); i++ {
			n2 := append([]int{}, n...)
			dfs(c, n2[i+1:], append(temp, n[i]))
		}
	}
	go func() {
		for i := 1; i <= len(nums)/2; i++ {
			dfs(i, nums, nil)
		}
		close(ch)
	}()
	return func(stop bool) ([]int, bool) {
		f, c := <-ch
		if stop {
			end = true
			return nil, false
		}
		return f, c
	}
}

func permute(nums []int) func() ([]int, bool) {
	ch := make(chan []int)
	var dfs func(l []int, temp []int)
	dfs = func(l []int, temp []int) {
		if len(l) == 0 {
			ch <- append([]int{}, temp...)
		}
		for i := 0; i < len(l); i++ {
			n := append([]int{}, l...)
			dfs(append(n[:i], n[i+1:]...), append(temp, l[i]))
		}
	}
	go func() {
		dfs(nums, []int{})
		close(ch)
	}()
	return func() ([]int, bool) {
		f, c := <-ch
		return f, c
	}
}

func avg(nums []int) float64 {
	sum := sum(nums)
	return float64(sum) / float64(len(nums))
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func splitArraySameAveragePlanC(nums []int) bool {
	s := sum(nums)
	l := len(nums)

	dp := make(map[int]map[int]bool)
	ad := func(i, j int, v bool) {
		n := dp[i]
		if n == nil {
			dp[i] = make(map[int]bool)
		}
		dp[i][j] = v
	}

	l2 := l / 2
	ad(0, 0, true)
	for _, n := range nums {
		for i := l2; i >= 1; i-- {
			for x := range dp[i-1] {
				c := x + n
				if c*l == s*i {
					return true
				}
				ad(i, c, true)
			}
		}
	}
	return false
}
