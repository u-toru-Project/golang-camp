package chapter05

import "fmt"

// BitSwapRequired counts 1-bits in the XOR by shifting.
func BitSwapRequired(first, second int) int {
	count := 0
	for bits := uint32(int32(first) ^ int32(second)); bits != 0; bits >>= 1 {
		count += int(bits & 1)
	}
	return count
}

// BitSwapRequired2 clears the lowest set bit each step (Kernighan).
func BitSwapRequired2(first, second int) int {
	count := 0
	for bits := int32(first) ^ int32(second); bits != 0; bits &= bits - 1 {
		count++
	}
	return count
}

func RunQ506() {
	const first, second = 29, 15
	fmt.Printf("%d: %s\n", first, ToFullBinaryString(first))
	fmt.Printf("%d: %s\n", second, ToFullBinaryString(second))
	fmt.Printf("Required bits: %d / %d\n", BitSwapRequired(first, second), BitSwapRequired2(first, second))
}
