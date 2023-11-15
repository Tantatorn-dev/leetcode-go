package easy

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func IsPalindrome(s string) bool {
	ret1 := ""
	ret2 := ""

	s = strings.ToLower(s)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")

	for i, h := range s {
		t := s[len(s)-1-i]

		if string(h) != " " {
			ret1 += string(h)
		}

		if string(t) != " " {
			ret2 += string(t)
		}
	}

	return ret1 == ret2
}
