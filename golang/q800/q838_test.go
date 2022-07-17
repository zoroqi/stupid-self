package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestPushDominoes(t *testing.T) {
	stupid_self.AssertEqual(t, pushDominoesPlanB("..L."), "LLL.")
	stupid_self.AssertEqual(t, pushDominoesPlanB("...."), "....")
	stupid_self.AssertEqual(t, pushDominoesPlanB("LL"), "LL")
	stupid_self.AssertEqual(t, pushDominoesPlanB(".R.."), ".RRR")
	stupid_self.AssertEqual(t, pushDominoesPlanB("R..L"), "RRLL")
	stupid_self.AssertEqual(t, pushDominoesPlanB("L..R"), "L..R")
	stupid_self.AssertEqual(t, pushDominoesPlanB(".L.R...LR..L.."), "LL.RR.LLRRLL..")
	stupid_self.AssertEqual(t, pushDominoesPlanB(".......R...L.."), ".......RR.LL..")
	stupid_self.AssertEqual(t, pushDominoesPlanB("R..........L.."), "RRRRRRLLLLLL..")
	stupid_self.AssertEqual(t, pushDominoesPlanB("R...........L.."), "RRRRRR.LLLLLL..")
}
