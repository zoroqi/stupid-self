package q800

func eventualSafeNodesPlanA(graph [][]int) []int {
	l := len(graph)
	safeNode := make([]int8, l)
	var isRing func(int) bool
	isRing = func(n int) bool {
		if safeNode[n] > 0 {
			return safeNode[n] == 1
		}
		safeNode[n] = 1
		for _, v := range graph[n] {
			if isRing(v) {
				return true
			}
		}
		safeNode[n] = 2
		return false
	}
	r := []int{}
	for i := 0; i < l; i++ {
		if !isRing(i) {
			r = append(r, i)
		}
	}
	return r
}

//func eventualSafeNodesPlanA(graph [][]int) []int {
//	unsafeNode := map[int]int{}
//
//	var isRing func(int, map[int]bool) bool
//	allLoop := 0
//
//	isRing = func(n int, ringFlag map[int]bool) bool {
//		allLoop++
//		if unsafeNode[n] == 1 {
//			return true
//		}
//		if unsafeNode[n] == 2 {
//			return false
//		}
//		for _, v := range graph[n] {
//			if ringFlag[v] {
//				return true
//			}
//			ringFlag[v] = true
//			n := isRing(v, ringFlag)
//			if !n {
//				ringFlag[v] = false
//			} else {
//				return true
//			}
//		}
//		return false
//	}
//	l := len(graph)
//	n := 0
//	bs, bl := 0, 0
//	for i := 0; i < l; i++ {
//		if unsafeNode[i] != 0 {
//			continue
//		}
//		fmt.Println(n, i, len(unsafeNode), "(", len(unsafeNode)-bs, ")", statis(unsafeNode), allLoop, "(", allLoop-bl, ")")
//		bs = len(unsafeNode)
//		bl = allLoop
//		n++
//		flags := map[int]bool{}
//		ring := isRing(i, flags)
//		if ring {
//			unsafeNode[i] = 1
//			for k, v := range flags {
//				if v {
//					unsafeNode[k] = 1
//				}
//			}
//		} else {
//			unsafeNode[i] = 2
//			for k := range flags {
//				unsafeNode[k] = 2
//			}
//		}
//	}
//	r := []int{}
//	for i := 0; i < l; i++ {
//		if unsafeNode[i] != 1 {
//			r = append(r, i)
//		}
//	}
//	fmt.Println(n, allLoop)
//	return r
//}
