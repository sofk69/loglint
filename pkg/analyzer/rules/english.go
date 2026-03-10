package rules

func IsEnglish(msg string) bool {
	for _, r := range msg {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == ' ' || r == '.' || r == ',') {
			return false
		}
	}
	return true
}
