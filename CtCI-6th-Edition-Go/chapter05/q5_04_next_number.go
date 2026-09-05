package chapter05

import "fmt"

func ToFullBinaryString(number int) string {
	n := int32(number)
	buf := make([]byte, 32)
	for i := 31; i >= 0; i-- {
		buf[i] = '0' + byte(n&1)
		n >>= 1
	}
	return string(buf)
}

// CountOnes counts 1-bits in the 32-bit representation.
func CountOnes(number int) int {
	count := 0
	for bits := uint32(int32(number)); bits != 0; bits >>= 1 {
		count += int(bits & 1)
	}
	return count
}

func CountZeros(number int) int {
	return 32 - CountOnes(number)
}

func HasValidNext(number int) bool {
	if number == 0 {
		return false
	}
	bits := uint32(int32(number))
	count := 0
	for bits&1 == 0 {
		bits >>= 1
		count++
	}
	for bits&1 == 1 {
		bits >>= 1
		count++
	}
	return count < 31
}

func HasValidPrev(number int) bool {
	bits := uint32(int32(number))
	for bits&1 == 1 {
		bits >>= 1
	}
	return bits != 0
}

func GetNextSlow(number int) int {
	if !HasValidNext(number) {
		return -1
	}
	ones := CountOnes(number)
	n := int32(number) + 1
	for CountOnes(int(n)) != ones {
		n++
	}
	return int(n)
}

func GetPrevSlow(number int) int {
	if !HasValidPrev(number) {
		return -1
	}
	ones := CountOnes(number)
	n := int32(number) - 1
	for CountOnes(int(n)) != ones {
		n--
	}
	return int(n)
}

func GetNext(number int) int {
	bits := uint32(int32(number))
	trailingZeros := 0
	trailingOnes := 0
	for bits&1 == 0 && bits != 0 {
		trailingZeros++
		bits >>= 1
	}
	for bits&1 == 1 {
		trailingOnes++
		bits >>= 1
	}
	if trailingZeros+trailingOnes == 0 || trailingZeros+trailingOnes >= 31 {
		return -1
	}
	n := int32(number)
	pos := trailingZeros + trailingOnes
	n |= 1 << pos
	n &= ^((1 << pos) - 1)
	n |= (1 << (trailingOnes - 1)) - 1
	return int(n)
}

func GetNextArith(number int) int {
	bits := uint32(int32(number))
	trailingZeros := 0
	trailingOnes := 0
	for bits&1 == 0 && bits != 0 {
		trailingZeros++
		bits >>= 1
	}
	for bits&1 == 1 {
		trailingOnes++
		bits >>= 1
	}
	if trailingZeros+trailingOnes == 0 || trailingZeros+trailingOnes >= 31 {
		return -1
	}
	n := int32(number) + (1 << trailingZeros) + (1 << (trailingOnes - 1)) - 1
	return int(n)
}

func GetPrev(number int) int {
	bits := uint32(int32(number))
	trailingZeros := 0
	trailingOnes := 0
	for bits&1 == 1 {
		trailingOnes++
		bits >>= 1
	}
	if bits == 0 {
		return -1
	}
	for bits&1 == 0 {
		trailingZeros++
		bits >>= 1
	}
	pos := trailingZeros + trailingOnes
	n := int32(number)
	if pos >= 31 {
		n = 0
	} else {
		n &= ^int32(0) << (pos + 1)
	}
	mask := int32((1 << (trailingOnes + 1)) - 1)
	n |= mask << (trailingZeros - 1)
	return int(n)
}

func GetPrevArith(number int) int {
	bits := uint32(int32(number))
	trailingZeros := 0
	trailingOnes := 0
	for bits&1 == 1 && bits != 0 {
		trailingOnes++
		bits >>= 1
	}
	if bits == 0 {
		return -1
	}
	for bits&1 == 0 && bits != 0 {
		trailingZeros++
		bits >>= 1
	}
	n := int32(number) - (1 << trailingOnes) - (1 << (trailingZeros - 1)) + 1
	return int(n)
}

func BinPrint(number int) {
	fmt.Printf("%d: %s\n", number, ToFullBinaryString(number))
}

func RunQ504() {
	for _, number := range []int{13948, 1, 2, 3, 6, 8} {
		fmt.Printf("%d (%s): prev %d / next %d\n", number, ToFullBinaryString(number), GetPrev(number), GetNext(number))
	}
}
