package rules

import "testing"

func TestIsEnglish(t *testing.T) {

	if IsEnglish("запуск сервера") {
		t.Fatal("expected false for russian text")
	}

	if !IsEnglish("starting server") {
		t.Fatal("expected true for english text")
	}

}
