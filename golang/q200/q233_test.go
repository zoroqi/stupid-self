package q200

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestCountDigitOnePlanA(t *testing.T) {
	stupid_self.AssertEqual(t, countDigitOnePlanA(13), 6)
	stupid_self.AssertEqual(t, countDigitOnePlanA(0), 0)
	stupid_self.AssertEqual(t, countDigitOnePlanA(15), 8)
	stupid_self.AssertEqual(t, countDigitOnePlanA(99), 20)
	stupid_self.AssertEqual(t, countDigitOnePlanA(832), 274)
	stupid_self.AssertEqual(t, countDigitOnePlanA(8305), 3561)
	stupid_self.AssertEqual(t, countDigitOnePlanA(8315), 3568)
	stupid_self.AssertEqual(t, countDigitOnePlanA(8310), 3562)
	stupid_self.AssertEqual(t, countDigitOnePlanA(8000), 3400)
}

func TestCountDigitOne(t *testing.T) {
	randsource := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 200; i++ {
		n := int(math.Abs(float64(randsource.Int()%50000) + 1))
		a := countDigitOnePlanA(n)
		b := countDigitOnePlanB(n)
		if a != b {
			t.Errorf("n=%d, a=%d, b=%d, sub=%d", n, a, b, a-b)
		}
	}
}
