package functions

import "testing"

func TestNonPrintable(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{"tab",
			"hello\\tworld",
			"hello    world",
		},
		{"backspace",
			"hello\\bworld",
			"hello world",
		},
		{"newline",
			"hello\\nworld",
			"hello\nworld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonPrintable(tt.args); got != tt.want {
				t.Errorf("NonPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}
