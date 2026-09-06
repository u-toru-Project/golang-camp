package chapter01

import "fmt"

func ReplaceSpaces(input []byte, trueLength int) {
	spaceCount := 0
	for i := range trueLength {
		if input[i] == ' ' {
			spaceCount++
		}
	}
	writeIndex := trueLength + spaceCount * 2 - 1
	for readIndex := trueLength - 1; readIndex >= 0;
	readIndex-- {
		if input[readIndex] == ' ' {
			input[writeIndex] = '0'
			writeIndex--
			input[writeIndex] = '2'
			writeIndex--
			input[writeIndex] = '%'
			writeIndex--
		} else {
			input[writeIndex] = input[readIndex]
			writeIndex--
		}
	}
}

func CreateBuffer(input string) []byte {
	spaceCount := 0
	for _, character := range input {
		if character == ' ' {
			spaceCount++
		}
	}
	buf := make([]byte, len(input) + spaceCount * 2)
	copy(buf, input)
	return buf
}

func RunQ103() {
	input := "abc d e f"
	buf := CreateBuffer(input)
	ReplaceSpaces(buf, len(input))
	fmt.Printf("%s -> %s\n", input, string(buf))
}

