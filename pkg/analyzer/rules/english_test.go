package rules

import "testing"

func TestIsEnglish(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "russian text", msg: "запуск сервера", want: false},
		{name: "mixed russian and english", msg: "server запущен", want: false},
		{name: "chinese characters", msg: "server 启动", want: false},
		{name: "simple english", msg: "starting server", want: true},
		{name: "english with numbers", msg: "retry attempt 3", want: true},
		{name: "empty string", msg: "", want: true},
		{name: "single word", msg: "connected", want: true},
		{name: "dot is not allowed", msg: "server started.", want: false},
		{name: "comma is not allowed", msg: "failed, retrying", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEnglish(tt.msg)
			if got != tt.want {
				t.Errorf("IsEnglish(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
