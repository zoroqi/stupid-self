package q0

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	stupid_self.AssertEqual(t, SimplifyPath("/home//foo/"), "/home/foo")
	stupid_self.AssertEqual(t, SimplifyPath("/home"), "/home")
	stupid_self.AssertEqual(t, SimplifyPath("/.."), "/")
	stupid_self.AssertEqual(t, SimplifyPath("/a/./b/../../c/"), "/c")
	stupid_self.AssertEqual(t, SimplifyPath("/a/../../b/../c//.//"), "/c")
	stupid_self.AssertEqual(t, SimplifyPath("/a//b////c/d//././/.."), "/a/b/c")
	stupid_self.AssertEqual(t, SimplifyPath(""), "/")
	stupid_self.AssertEqual(t, SimplifyPath("/"), "/")
	stupid_self.AssertEqual(t, SimplifyPath("/a"), "/a")
}
