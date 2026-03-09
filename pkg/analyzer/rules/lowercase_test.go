package rules

import "testing"

func TestIsLowercase(t *testing.T) {

	if IsLowercase("Starting server") {
		t.Fatal("should be false")
	}

	if !IsLowercase("starting server") {
		t.Fatal("should be true")
	}

}
