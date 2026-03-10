package rules

import "testing"

func TestIsLowercase(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "uppercase first letter", msg: "Starting server", want: false},
		{name: "all uppercase", msg: "STARTING SERVER", want: false},
		{name: "uppercase single word", msg: "Failed", want: false},
		{name: "lowercase first letter", msg: "starting server", want: true},
		{name: "lowercase single word", msg: "connected", want: true},
		{name: "starts with digit", msg: "3 retries left", want: true},
		{name: "russian lowercase", msg: "запуск сервера", want: true},
		{name: "empty string", msg: "", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLowercase(tt.msg)
			if got != tt.want {
				t.Errorf("IsLowercase(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
