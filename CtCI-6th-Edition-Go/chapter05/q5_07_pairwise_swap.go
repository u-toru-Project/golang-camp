package chapter05

import "fmt"

// SwapOddEvenBits swaps odd and even bits using a logical shift.
func SwapOddEvenBits(x int) int {
	bits := uint32(int32(x))
	return int(int32(((bits & 0xaaaaaaaa) >> 1) | ((bits & 0x55555555) << 1)))
}

func RunQ507() {
	const a = 103217
	swapped := SwapOddEvenBits(a)
	fmt.Printf("%d: %s\n", a, ToFullBinaryString(a))
	fmt.Printf("%d: %s\n", swapped, ToFullBinaryString(swapped))
}
