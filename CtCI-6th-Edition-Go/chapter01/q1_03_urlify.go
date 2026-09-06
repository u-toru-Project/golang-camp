package chapter01

import "fmt"

const (
	spaceRune     = ' '
	urlSpace      = "%20"
	urlSpaceLen   = len(urlSpace)
	extraPerSpace = urlSpaceLen - 1 // ' ' (1 byte) expands to "%20" (3 bytes)
)

func countSpacesInString(input string) int {
	count := 0
	for _, character := range input {
		if character == spaceRune {
			count++
		}
	}
	return count
}

func countSpacesInPrefix(input []byte, length int) int {
	count := 0
	for i := range length {
		if input[i] == spaceRune {
			count++
		}
	}
	return count
}

func extraLengthForSpaces(spaceCount int) int {
	return spaceCount * extraPerSpace
}

// ReplaceSpaces writes %20 into input in place. trueLength is the meaningful prefix.
// Time O(n), space O(1).
func ReplaceSpaces(input []byte, trueLength int) {
	spaceCount := countSpacesInPrefix(input, trueLength)
	writeIndex := trueLength + extraLengthForSpaces(spaceCount) - 1
	for readIndex := trueLength - 1; readIndex >= 0; readIndex-- {
		if input[readIndex] == spaceRune {
			writeIndex -= urlSpaceLen - 1
			copy(input[writeIndex:writeIndex+urlSpaceLen], urlSpace)
			writeIndex--
		} else {
			input[writeIndex] = input[readIndex]
			writeIndex--
		}
	}
}

func CreateBuffer(input string) []byte {
	spaceCount := countSpacesInString(input)
	buf := make([]byte, len(input)+extraLengthForSpaces(spaceCount))
	copy(buf, input)
	return buf
}

func RunQ103() {
	input := "abc d e f"
	buf := CreateBuffer(input)
	ReplaceSpaces(buf, len(input))
	fmt.Printf("%s -> %s\n", input, string(buf))
}
