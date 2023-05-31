package q600

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func buildSpecial662Tree() *stupid_self.TreeNode {
	left, right := &stupid_self.TreeNode{}, &stupid_self.TreeNode{}
	root := &stupid_self.TreeNode{Left: left, Right: right}
	for i := 0; i < 25; i++ {
		left.Right = &stupid_self.TreeNode{}
		right.Left = &stupid_self.TreeNode{}
		left = left.Right
		right = right.Left
	}
	return root
}
func Test_widthOfBinaryTree(t *testing.T) {
	stupid_self.AssertEqual(t, widthOfBinaryTreePlanB(stupid_self.NewTreeNode([]int{1, 3, 2, 5, 3, 0, 9}, 0)), 4)
	stupid_self.AssertEqual(t, widthOfBinaryTreePlanB(stupid_self.NewTreeNode([]int{1, 3, 2, 5, 0, 0, 9, 6, 0, 0, 0, 0, 0, 7}, 0)), 7)
	stupid_self.AssertEqual(t, widthOfBinaryTreePlanB(stupid_self.NewTreeNode([]int{1, 3, 2, 5}, 0)), 2)
	stupid_self.AssertEqual(t, widthOfBinaryTreePlanB(buildSpecial662Tree()), 2)
}

func Benchmark_widthOfBinaryTree(b *testing.B) {
	root := buildSpecial662Tree()
	for i := 0; i < b.N; i++ {
		widthOfBinaryTreePlanA(root)
	}
}
