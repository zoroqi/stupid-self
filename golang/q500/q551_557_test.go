package q500

import (
	"fmt"
	"testing"
)

func TestCheckRecord(t *testing.T) {
	fmt.Println(CheckRecord("PPALLL"))
}

func TestReverseWords(t *testing.T) {
	s := " Let's take d  LeetCode contest    "
	fmt.Println(s)
	fmt.Println(ReverseWords(s))
}