package rules

import "unicode"

func HasSpecialChars(msg string) bool {
	for _, r := range msg {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == ' ':
			continue

		case unicode.Is(unicode.So, r):
			return true
		case unicode.Is(unicode.Sk, r):
			return true
		case r > 0x2000:
			return true

		case unicode.IsPunct(r):
			return true
		case unicode.IsSymbol(r):
			return true
		}
	}
	return false
}
