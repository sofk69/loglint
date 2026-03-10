package rules

import "testing"

func TestHasSpecialChars(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "exclamation mark", msg: "server started!", want: true},
		{name: "multiple exclamations", msg: "connection failed!!!", want: true},
		{name: "ellipsis", msg: "something went wrong...", want: true},
		{name: "colon", msg: "warning: something wrong", want: true},
		{name: "rocket emoji", msg: "server started 🚀", want: true},
		{name: "fire emoji", msg: "deployed 🔥", want: true},
		{name: "checkmark emoji", msg: "done ✅", want: true},
		{name: "clean message", msg: "server started", want: false},
		{name: "message with numbers", msg: "retry attempt 3", want: false},
		{name: "message with hyphen", msg: "connection failed", want: false},
		{name: "empty string", msg: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasSpecialChars(tt.msg)
			if got != tt.want {
				t.Errorf("HasSpecialChars(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
