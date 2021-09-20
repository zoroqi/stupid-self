package util

import (
	"fmt"
	"testing"
)

func TestBytesPrefix(t *testing.T) {
	r := BytesPrefix([]byte("abcdef"))
	fmt.Println(r.Start,r.Limit)
}