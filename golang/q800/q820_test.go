package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestMinimumLengthEncoding(t *testing.T) {
	stupid_self.AssertEqual(t, minimumLengthEncoding([]string{"time", "me", "bell"}), 10)
	stupid_self.AssertEqual(t, minimumLengthEncoding([]string{"t"}), 2)
	stupid_self.AssertEqual(t, minimumLengthEncoding([]string{"feipyxx", "e"}), 10)
	stupid_self.AssertEqual(t, minimumLengthEncoding([]string{"grah", "p", "qwosp"}), 11)
	stupid_self.AssertEqual(t, minimumLengthEncoding([]string{"qaa", "a", "aa"}), 4)
}

func TestMinimumLengthEncodingPlanB(t *testing.T) {
	stupid_self.AssertEqual(t, minimumLengthEncodingPlanB([]string{"time", "me", "bell"}), 10)
	stupid_self.AssertEqual(t, minimumLengthEncodingPlanB([]string{"t"}), 2)
	stupid_self.AssertEqual(t, minimumLengthEncodingPlanB([]string{"feipyxx", "e"}), 10)
	stupid_self.AssertEqual(t, minimumLengthEncodingPlanB([]string{"grah", "p", "qwosp"}), 11)
	stupid_self.AssertEqual(t, minimumLengthEncodingPlanB([]string{"qaa", "a", "aa"}), 4)
}
