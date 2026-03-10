package rules

import "testing"

func TestContainsSensitive(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "password keyword", msg: "user password: secret123", want: true},
		{name: "passwd keyword", msg: "passwd value loaded", want: true},
		{name: "token keyword", msg: "token: abc123", want: true},
		{name: "api_key keyword", msg: "api_key= xyz", want: true},
		{name: "apikey keyword", msg: "apikey loaded", want: true},
		{name: "secret keyword", msg: "secret value loaded", want: true},
		{name: "private_key keyword", msg: "private_key found", want: true},
		{name: "credential keyword", msg: "credential expired", want: true},
		{name: "uppercase password", msg: "User PASSWORD is set", want: true},
		{name: "mixed case token", msg: "Bearer TOKEN found", want: true},
		{name: "authenticated is clean", msg: "user authenticated successfully", want: false},
		{name: "clean api message", msg: "api request completed", want: false},
		{name: "session validated", msg: "session validated", want: false},
		{name: "authorization header", msg: "authorization header missing", want: false},
		{name: "empty string", msg: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsSensitive(tt.msg)
			if got != tt.want {
				t.Errorf("ContainsSensitive(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
