package rules

import "strings"

var forbidden = []string{
	"!",
	"...",
	"🚀",
}

func HasSpecialChars(msg string) bool {

	for _, f := range forbidden {

		if strings.Contains(msg, f) {
			return true
		}

	}

	return false
}
