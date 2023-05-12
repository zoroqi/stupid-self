package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestPushDominoes(t *testing.T) {

	type td struct {
		dominoes string
		r        string
	}
	tds := []td{
		{"..L.", "LLL."},
		{"....", "...."},
		{"LL", "LL"},
		{".R..", ".RRR"},
		{"R..L", "RRLL"},
		{"L..R", "L..R"},
		{".L.R...LR..L..", "LL.RR.LLRRLL.."},
		{".......R...L..", ".......RR.LL.."},
		{"R..........L..", "RRRRRRLLLLLL.."},
		{"R...........L..", "RRRRRR.LLLLLL.."},
	}
	for _, v := range tds {
		stupid_self.AssertEqual(t, pushDominoesPlanA(v.dominoes), v.r)
	}
	for _, v := range tds {
		stupid_self.AssertEqual(t, pushDominoesPlanB(v.dominoes), v.r)
	}
}
