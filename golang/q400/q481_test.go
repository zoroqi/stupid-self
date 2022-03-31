package q400

import (
	"fmt"
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestPrintProgression(t *testing.T) {
	o := levelProgression(5)
	n := infiniteProgression(400)
	for i := 0; i < len(o); i++ {
		if o[i] != n[i] {
			fmt.Println(i, o[i], n[i])
			for j := 0; j <= i; j++ {
				fmt.Println(j, o[j], n[j])
			}
			break
		}
	}
}

func TestMagicalString(t *testing.T) {
	stupid_self.AssertEqual(t, magicalString(6), 3)
	stupid_self.AssertEqual(t, magicalString(3), 1)
	stupid_self.AssertEqual(t, magicalString(11), 5)
	stupid_self.AssertEqual(t, magicalString(22029), 11007)
}
