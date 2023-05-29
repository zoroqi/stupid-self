package q1900

type LockingTree struct {
	tree   []*q1933node
	locker map[int]int
	length int
}

type q1933node struct {
	parent   int
	selfId   int
	children []*q1933node
}

func buildQ1933Tree(parent []int) []*q1933node {
	length := len(parent)
	nodes := make([]*q1933node, length)
	for i := 0; i < length; i++ {
		nodes[i] = &q1933node{parent: parent[i], selfId: i}
	}
	for i := 0; i < length; i++ {
		if parent[i] >= 0 {
			nodes[parent[i]].children = append(nodes[parent[i]].children, nodes[i])
		}
	}
	return nodes
}

func Q1933Constructor(parent []int) LockingTree {
	return LockingTree{buildQ1933Tree(parent), make(map[int]int, len(parent)), len(parent)}
}

func (this *LockingTree) unlock(num int) bool {
	if n := this.locker[num]; n > 0 {
		this.locker[num] = -1
		return true
	}
	return false
}

func (this *LockingTree) isLock(num int) bool {
	return this.locker[num] > 0
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.locker[num] > 0 {
		return false
	}
	this.locker[num] = user
	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.locker[num] != user {
		return false
	}
	return this.unlock(num)
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	// 自身和父节点没有被加锁
	for n := num; n >= 0; {
		if this.isLock(n) {
			return false
		}
		n = this.tree[n].parent
	}
	// 所有子节点
	length := this.length
	lockChildren := []int{}
	tree := this.tree
	var dfs func(int)
	dfs = func(i int) {
		if i >= length {
			return
		}
		if this.isLock(i) {
			lockChildren = append(lockChildren, i)
		}
		node := tree[i]
		if len(node.children) > 0 {
			for _, c := range node.children {
				dfs(c.selfId)
			}
		}
	}
	dfs(num)
	if len(lockChildren) > 0 {
		for _, i := range lockChildren {
			this.unlock(i)
		}
		this.Lock(num, user)
		return true
	}
	return false
}
