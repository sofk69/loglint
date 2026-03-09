package rules

func IsEnglish(msg string) bool {

	for _, r := range msg {

		if r > 127 {
			return false
		}

	}

	return true
}
