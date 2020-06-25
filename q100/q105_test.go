package q100

import (
	stupid_self "github.com/zoroqi/stupid-self"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	inorder := []int{4, 2, 5, 1, 6, 3, 7}
	node := BuildTree(preorder, inorder)
	stupid_self.PrintTreeNode(node)
}

func TestBuildTree2(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	node := BuildTree(preorder, inorder)
	stupid_self.PrintTreeNode(node)
}
func TestPreAndPostBuildTree(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	node := PreAndPostBuildTree(preorder, postorder)
	stupid_self.PrintTreeNode(node)
}
