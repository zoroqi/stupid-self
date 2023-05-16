package q1000

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestCamelMatchx(t *testing.T) {
	stupid_self.AssertEqual(t,
		camelMatchPlanA([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"),
		[]bool{false, true, false, false, false})
	stupid_self.AssertEqual(t,
		camelMatchPlanA([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"),
		[]bool{true, false, true, true, false})
	stupid_self.AssertEqual(t,
		camelMatchPlanA([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"),
		[]bool{true, false, true, false, false})
	stupid_self.AssertEqual(t,
		camelMatchPlanA([]string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}, "CooP"),
		[]bool{false, false, true})
	stupid_self.AssertEqual(t,
		camelMatchPlanA([]string{"aksvbjLiknuTzqon", "ksvjLimflkpnTzqn", "mmkasvjLiknTxzqn", "ksvjLiurknTzzqbn", "ksvsjLctikgnTzqn", "knzsvzjLiknTszqn"}, "ksvjLiknTzqn"),
		[]bool{true, true, true, true, true, true})
}
