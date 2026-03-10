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

var extraSensitive []string

func SetExtraSensitive(keywords []string) {
	extraSensitive = keywords
}

func ContainsSensitive(msg string) bool {
	msg = strings.ToLower(msg)
	all := append(sensitive, extraSensitive...)
	for _, s := range all {
		if strings.Contains(msg, s) {
			return true
		}
	}
	return false
}
