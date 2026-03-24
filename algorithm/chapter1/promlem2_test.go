package chapter1

import (
	"testing"
)

func TestArePermutations(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		want   bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abdc", true},
		{"abcc", "ccbb", false},
		{"abcc", "abcc ", false},
		{" ", " ", true},
		{"", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.input1+"/"+tt.input2,
			func(t *testing.T) {
				if got := ArePermutations(tt.input1, tt.input2); got != tt.want {
					t.Errorf("ArePermtations(%q, %q) = %t, ant %t", tt.input1, tt.input2, got, tt.want)
				}
			})
	}
}
