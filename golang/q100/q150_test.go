package q100

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	t.Log(EvalRPN([]string{"2", "1", "+", "3", "*"}))
	t.Log(EvalRPN([]string{"4", "13", "5", "/", "+"}))
}

var arr = []string{"22", "10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+", "*"}

func BenchmarkEvalRPN(b *testing.B) {
	b.Run(fmt.Sprintf("1"), func(b *testing.B) {
		EvalRPN(arr)
	})
}

func BenchmarkFirst150(b *testing.B) {
	b.Run(fmt.Sprintf("1"), func(b *testing.B) {
		first150(arr)
	})
}
