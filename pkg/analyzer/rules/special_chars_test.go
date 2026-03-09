package rules

import "testing"

func TestSpecialChars(t *testing.T) {

	if !HasSpecialChars("server started!") {
		t.Fatal("should detect special char")
	}

	if HasSpecialChars("server started") {
		t.Fatal("should not detect")
	}

}
