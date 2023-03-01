package q800

func possibleBipartitionPlanA(n int, dislikes [][]int) bool {
	if n <= 2 {
		return true
	}
	dislikeMap := map[int]map[int]bool{}
	for _, dis := range dislikes {
		if _, e := dislikeMap[dis[0]]; !e {
			dislikeMap[dis[0]] = map[int]bool{}
		}
		if _, e := dislikeMap[dis[1]]; !e {
			dislikeMap[dis[1]] = map[int]bool{}
		}
		dislikeMap[dis[0]][dis[1]] = true
		dislikeMap[dis[1]][dis[0]] = true
	}
	g1 := map[int]bool{}
	g2 := map[int]bool{}
	noDislike := map[int]bool{}
	for i := 1; i <= n; i++ {
		if dislike, exist := dislikeMap[i]; exist {
			if g2[i] {
				g1, g2 = g2, g1
			}
			g1[i] = true
			for d := range dislike {
				if g1[d] {
					return false
				}
				if d2, e := dislikeMap[d]; e {
					for d2d := range d2 {
						if g2[d2d] {
							return false
						}
					}
				}
				g2[d] = true
			}
		} else {
			noDislike[i] = true
		}
	}
	return true
}

func possibleBipartitionPlanB(n int, dislikes [][]int) bool {
	if n <= 2 {
		return true
	}
	dislikeMap := map[int]map[int]bool{}
	for _, dis := range dislikes {
		if _, e := dislikeMap[dis[0]]; !e {
			dislikeMap[dis[0]] = map[int]bool{}
		}
		if _, e := dislikeMap[dis[1]]; !e {
			dislikeMap[dis[1]] = map[int]bool{}
		}
		dislikeMap[dis[0]][dis[1]] = true
		dislikeMap[dis[1]][dis[0]] = true
	}
	var dfs func(i int, g1, g2 map[int]bool) bool
	dfs = func(i int, g1, g2 map[int]bool) bool {
		if g1[i] || g2[i] {
			return true
		}
		if dislike, exist := dislikeMap[i]; exist {
			if g2[i] {
				g1, g2 = g2, g1
			}
			g1[i] = true
			for d := range dislike {
				if g1[d] {
					return false
				}
				if d2, e := dislikeMap[d]; e {
					for d2d := range d2 {
						if g2[d2d] {
							return false
						}
					}
				}
				g2[d] = true
				if d2, e := dislikeMap[d]; e {
					for d2d := range d2 {
						r := dfs(d2d, g1, g2)
						if !r {
							return r
						}
					}
				}
			}
		}
		return true
	}
	g1, g2 := map[int]bool{}, map[int]bool{}
	for i := 1; i < n; i++ {
		if g1[i] || g2[i] {
			continue
		}
		r := dfs(i, g1, g2)
		if !r {
			return r
		}
	}
	return true
}

func possibleBipartitionPlanC(n int, dislikes [][]int) bool {
	if n <= 2 || len(dislikes) <= 1 {
		return true
	}
	dislikeGraph := make([][]int, n)
	for _, dis := range dislikes {
		d0 := dis[0] - 1
		d1 := dis[1] - 1
		dislikeGraph[d0] = append(dislikeGraph[d0], d1)
		dislikeGraph[d1] = append(dislikeGraph[d1], d0)
	}

	var dfs func(i int, g1, g2 map[int]bool) bool
	dfs = func(i int, g1, g2 map[int]bool) bool {
		if g1[i] || g2[i] {
			return true
		}

		g1[i] = true
		for _, d := range dislikeGraph[i] {
			if g1[d] {
				return false
			}
			if !dfs(d, g2, g1) {
				return false
			}
		}
		return true
	}
	g1, g2 := map[int]bool{}, map[int]bool{}
	for i := 0; i < n; i++ {
		if !dfs(i, g1, g2) {
			return false
		}
	}
	return true
}

func possibleBipartitionFastest(n int, dislikes [][]int) bool {
	if n <= 2 || len(dislikes) <= 1 {
		return true
	}
	// 表示的分组 3 未分组；1 第一组； 2 第二组
	// 为什么选择3？
	//	因为：3-1=2 3-2=1
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = 3
	}
	// 记录每个节点不能一组的节点
	hash := make([][]int, n)
	for _, v := range dislikes {
		var a, b = v[0] - 1, v[1] - 1
		hash[a] = append(hash[a], b)
		hash[b] = append(hash[b], a)
	}

	var dfs func(i, v int) bool
	dfs = func(i, v int) bool {
		l[i] = v
		for _, x := range hash[i] {
			if l[x] == v || l[x] == 3 && !dfs(x, 3-l[i]) {
				return false
			}
		}
		return true
	}
	for i, v := range l {
		if v == 3 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

func possibleBipartitionPlanD(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	uf := newUnionFind(n)
	for x, nodes := range g {
		for _, y := range nodes {
			// 将整个 nodes 合并到一个树内
			uf.union(nodes[0], y)
			// 判断 x 是否和 y 在同一个树内
			if uf.isConnected(x, y) {
				return false
			}
		}
	}
	return true
}

type q886unionFind struct {
	// parent[i] 表示第 i 个元素所指向的父节点,
	//parent[i] == i 表示第 i 个元素指向根节点,
	parent []int
	// rank[i] 表示以 i 为根的集合所表示的树的层数
	rank []int
}

func newUnionFind(n int) q886unionFind {
	// 初始化所有节点都是根节点, 之后进行联合.
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return q886unionFind{parent, make([]int, n)}
}

func (uf q886unionFind) find(x int) int {
	// 压缩路径, 递归的把所有路径的节点都挂到根节点上
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	// 返回根节点
	return uf.parent[x]
}

func (uf q886unionFind) union(x, y int) {
	x, y = uf.find(x), uf.find(y)
	// 如果 x 和 y 已经在同一个集合中, 则不需要合并
	if x == y {
		return
	}
	// 将层数小的合并到层数大的, 保证层数最大的树的层数不会超过 logN
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}

func (uf q886unionFind) isConnected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}
