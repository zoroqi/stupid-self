package q800

import (
	stupid_self "github.com/zoroqi/stupid-self/golang"
	"testing"
)

func TestMail831(t *testing.T) {
	stupid_self.AssertEqual(t, mail_831("AB@qq.com"), "a*****b@qq.com")
}

func TestPhone831(t *testing.T) {
	stupid_self.AssertEqual(t, phone_831("1(234)567-890"), "***-***-7890")
	stupid_self.AssertEqual(t, phone_831("86-(10)12345678"), "+**-***-***-5678")
	stupid_self.AssertEqual(t, phone_831("+(501321)-50-23431"), "+***-***-***-3431")
}
