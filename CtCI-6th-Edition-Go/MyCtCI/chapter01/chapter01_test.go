package chapter01

import (
	"testing"
)

func TestIsUniqueChars(t *testing.T) {
	if !IsUniqueChars("") || !IsUniqueChars("abcde") || IsUniqueChars("hello") {
		t.Fatal("IsUniqueChars mismatch")
	}
	if _, err := IsUniqueAsciiLetters("ABC"); err == nil {
		t.Fatal("expected error")
	}
	ok, err := IsUniqueAsciiLetters("abc")
	if err != nil || !ok {
		t.Fatal(ok, err)
	}
}

func TestPermutation(t *testing.T) {
	if !IsPermutationBySorting("apple", "papel") || !IsPermutationByCounts("apple", "papel") {
		t.Fatal("expected permutation")
	}
	if IsPermutationBySorting("hello", "llloh") || IsPermutationByCounts("Ab", "ab") {
		t.Fatal("expected non-permutation")
	}
}
