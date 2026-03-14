package chapter1

import (
	"testing"
)

func TestIsUniqueCheck(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"重複なしのパターン", "dog", true},
		{"重複ありのパターン", "aba", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUniqueCheck(tt.input); got != tt.want {
				t.Errorf("IsUniqueCheck(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
