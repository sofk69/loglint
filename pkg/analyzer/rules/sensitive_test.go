package rules

import "testing"

func TestSensitive(t *testing.T) {

	if !ContainsSensitive("user password") {
		t.Fatal("should detect sensitive word")
	}

	if ContainsSensitive("user authenticated") {
		t.Fatal("false positive")
	}

}
