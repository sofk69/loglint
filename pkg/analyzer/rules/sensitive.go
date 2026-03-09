package rules

import "strings"

var sensitive = []string{
	"password",
	"token",
	"api_key",
}

func ContainsSensitive(msg string) bool {

	msg = strings.ToLower(msg)

	for _, s := range sensitive {

		if strings.Contains(msg, s) {
			return true
		}

	}

	return false
}
