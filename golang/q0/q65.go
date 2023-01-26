package q0

import (
	"regexp"
	"strings"
)

var reg *regexp.Regexp
var reg_e *regexp.Regexp

func init() {
	reg, _ = regexp.Compile("^[\\-\\+]?(\\d+)?(\\.\\d*?)?(e[\\-\\+]?\\d+)?$")
	reg_e, _ = regexp.Compile("^[\\.\\-\\+]?e[\\-\\+]?\\d+$")
}

func IsNumber(s string) bool {
	num := strings.TrimSpace(s)
	if len(num) == 0 {
		return false
	}
	if len(num) == 1 {
		return num[0] >= 48 && num[0] <= 57
	}
	if num == "-." || num == "+." {
		return false
	}
	return !reg_e.MatchString(num) && reg.MatchString(num)
}
