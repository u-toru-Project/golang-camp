package chapter01

import (
	"bytes"
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

func TestURLify(t *testing.T) {
	input := "Mr John Smith"
	buf := CreateBuffer(input)
	ReplaceSpaces(buf, len(input))
	if string(buf) != "Mr%20John%20Smith" {
		t.Fatalf("got %q", buf)
	}
	empty := CreateBuffer("")
	ReplaceSpaces(empty, 0)
	if !bytes.Equal(empty, []byte{}) && string(empty) != "" {
		t.Fatalf("empty: %q", empty)
	}
}

func TestPalindromeAndOneAway(t *testing.T) {
	if !IsPermutationOfPalindrome("Tact Coa") || !IsPermutationOfPalindrome2("Tact Coa") || !IsPermutationOfPalindrome3("Tact Coa") {
		t.Fatal("palindrome")
	}
	if IsPermutationOfPalindrome("asda") {
		t.Fatal("not palindrome")
	}
	if !OneEditAway("pale", "ple") || !OneEditAway2("pale", "bale") || OneEditAway("pale", "bake") {
		t.Fatal("one away")
	}
}

func TestCompressRotateZeroRotation(t *testing.T) {
	if Compress("aabcccccaaa") != "a2b1c5a3" || Compress("abc") != "abc" || Compress("aaaaaaaaaa") != "a10" {
		t.Fatal("compress")
	}
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if !Rotate(matrix) || matrix[0][0] != 7 || matrix[0][2] != 1 {
		t.Fatalf("rotate %#v", matrix)
	}
	if Rotate(nil) {
		t.Fatal("nil rotate")
	}
	zeros := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 0, 9}}
	SetZeros(zeros)
	if zeros[0][0] != 0 || zeros[1][0] != 4 || zeros[1][1] != 0 {
		t.Fatalf("zeros %#v", zeros)
	}
	if !IsRotation("waterbottle", "erbottlewat") || IsRotation("camera", "macera") || IsRotation("", "") {
		t.Fatal("rotation")
	}
}
