package rules

import "unicode"

func IsLowercase(msg string) bool {

	if len(msg) == 0 {
		return true
	}

	return unicode.IsLower(rune(msg[0]))
}
