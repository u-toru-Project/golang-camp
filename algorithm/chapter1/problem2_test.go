package chapter1

import (
	"testing"
)

func TestArePermutations(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{"same string", "abcd", "abcd", true},
		{"same characters in different order", "abcd", "abdc", true},
		{"different character counts", "abcc", "ccbb", false},
		{"different length because of trailing space", "abcc", "abcc ", false},
		{"single space", " ", " ", true},
		{"empty strings", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				if got := ArePermutations(tt.input1, tt.input2); got != tt.want {
					t.Errorf("ArePermutations(%q, %q) = %t, want %t", tt.input1, tt.input2, got, tt.want)
				}
			})
	}
}
