package rules

import "unicode"

func IsLowercase(msg string) bool {
	if len(msg) == 0 {
		return true
	}
	r := []rune(msg)[0]
	return !unicode.IsUpper(r)
}
