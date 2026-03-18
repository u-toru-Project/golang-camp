package chapter1

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcd", true},
		{"abcc", false},
		{" ", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsUnique(tt.input); got != tt.want {
				t.Errorf("IsUnique(%q) = %t, want %t", tt.input, got, tt.want)
			}
		})
	}
}
