package rules

import "strings"

var sensitive = []string{
	"password",
	"passwd",
	"token",
	"api_key",
	"apikey",
	"secret",
	"private_key",
	"credential",
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
