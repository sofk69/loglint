package rules

import "unicode"

func HasSpecialChars(msg string) bool {
	for _, r := range msg {
		if unicode.IsLetter(r) {
			if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
				return true
			}
			continue
		}
		if unicode.IsDigit(r) || r == ' ' {
			continue
		}
		return true
	}
	return false
}
