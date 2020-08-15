package q500

import "fmt"

func FindCircleNum(M [][]int) int {
	length := len(M)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	friends := make([]bool, length)
	var dfsFind func(int)
	dfsFind = func(j int) {
		for i := 0; i < length; i++ {
			if M[j][i] == 1 && !friends[i] {
				friends[i] = true
				dfsFind(i)
			}
		}
	}
	count := 0
	for i := 0; i < length; i++ {
		if !friends[i] {
			count++
			dfsFind(i)
		}
	}
	return count
}

func secondFindCircleNum(M [][]int) int {
	length := len(M)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}

	friends := make(map[int]map[int]bool)
	dfsFind := func(subchief int, MLM []int) []int {
		newMLM := make([]int, 0, length)
		for i, v := range MLM {
			if v == 1 {
				if friends[i] == nil {
					friends[i] = make(map[int]bool)
				}
				if !friends[subchief][i] {
					newMLM = append(newMLM, i)
				}
				friends[subchief][i] = true
			}
		}
		return newMLM
	}
	for i := 0; i < length; i++ {
		v := dfsFind(i, M[i])
		for len(v) > 0 {
			n := make([]int, 0)
			for _, v2 := range v {
				n = append(n, dfsFind(i, M[v2])...)
			}
			v = n
		}
		//fmt.Println("-------")
	}
	//printMap(friends)
	for k, v := range friends {
		for k2 := range v {
			if k != k2 {
				delete(friends, k2)
			}
		}
	}
	printMap(friends)
	return len(friends)
}
func firstFindCircleNum(M [][]int) int {
	length := len(M)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	m := make(map[int]map[int]bool, length)
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			if M[i][j] == 1 {
				if m[i] == nil {
					m[i] = make(map[int]bool)
				}
				m[i][j] = true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for k, v := range m {
			for k2 := range v {
				if m[k2] != nil {
					for k3 := range m[k2] {
						m[k3][k] = true
					}
				}
			}
		}
	}
	for k, v := range m {
		for k2 := range v {
			if k != k2 {
				delete(m, k2)
			}
		}
	}
	return len(m)
}

func printMap(m map[int]map[int]bool) {
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("--------")
}
