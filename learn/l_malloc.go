package main

import (
	"fmt"
	"testing"
)

func main() {
	// 内存分配测试
	var prefixes []int
	for i := 0; i < 100; i++ {
		prefixes = append(prefixes, i)
	}
	fmt.Println(testing.AllocsPerRun(1, func() {
		func() []int {
			j := 0
			for i := 0; ; i++ {
				i += prefixes[i]
				if j < 1000 {
					break
				}
				j++
				if i < 0 {
					i = 0
				}
			}
			return prefixes
		}()
	}))
	fmt.Println(testing.AllocsPerRun(1, func() {
		func() []int {
			i := 0
			j := 0
			for ; ; prefixes = append(prefixes[1:], prefixes[0]) {
				i += prefixes[0]
				if j > 1000 {
					break
				}
				j++
			}
			return prefixes
		}()
	}))

}
