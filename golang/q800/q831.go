package q800

import (
	"fmt"
	"regexp"
	"strings"
)

func maskPII(s string) string {
	if strings.Contains(s, "@") {
		return mail_831(s)
	} else {
		return phone_831(s)
	}

}

var mail_831_regex = regexp.MustCompile("(.)(.*?)(.)@(.+)")
var phone_831_regex = regexp.MustCompile("[{}()+\\- ]")

func mail_831(s string) string {
	pii := mail_831_regex.ReplaceAllString(s, "$1*****$3@$4")
	return strings.ToLower(pii)
}

func phone_831(s string) string {
	pii := phone_831_regex.ReplaceAllString(s, "")
	fmt.Println(pii)
	switch len(pii) {
	case 10:
		return "***-***-" + pii[len(pii)-4:]
	case 11:
		return "+*-***-***-" + pii[len(pii)-4:]
	case 12:
		return "+**-***-***-" + pii[len(pii)-4:]
	case 13:
		return "+***-***-***-" + pii[len(pii)-4:]
	}
	return s
}
