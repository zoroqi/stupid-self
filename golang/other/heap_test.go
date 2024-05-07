package other

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"math/rand/v2"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &Heap{}
	arr := []int{}
	for i := 1; i < 20; i++ {
		arr = append(arr, i)
		arr = append(arr, i)
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	var dfs func(node *stupid_self.TreeNode) bool
	dfs = func(node *stupid_self.TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Left == nil && node.Right == nil {
			return true
		}
		if node.Left != nil && node.Val > node.Left.Val {
			return false
		}
		if node.Right != nil && node.Val > node.Right.Val {
			return false
		}
		return dfs(node.Left) && dfs(node.Right)
	}
	for _, v := range arr {
		h.Push(v)
		if !dfs(stupid_self.NewTreeNode(h.queue[:h.size], 0)) {
			t.Fatal("push error")
		}
	}
	for range 10 {
		_, ok := h.Pop()
		if !ok {
			t.Fatal("top error")
		}
		if !dfs(stupid_self.NewTreeNode(h.queue[:h.size], 0)) {
			t.Fatal("heap error")
		}
	}
	for _, v := range arr {
		h.Push(v)
		if !dfs(stupid_self.NewTreeNode(h.queue[:h.size], 0)) {
			t.Fatal("push error")
		}
	}
	for range h.size {
		_, ok := h.Pop()
		if !ok {
			t.Fatal("top error")
		}
		if !dfs(stupid_self.NewTreeNode(h.queue[:h.size], 0)) {
			t.Fatal("heap error")
		}
	}
}
