package q500

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	fmt.Println(Fib(8))
	fmt.Println(Fib_haskell(8))
	fmt.Println(Fib_recursion(8))
}
