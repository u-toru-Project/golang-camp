package chapter05

import (
	"fmt"
	"math"
	"strconv"
)

type FloatDemo struct {
	DecimalStr     string
	BinaryFraction string
	Ieee754Hex     string
}

// IsPowerOfTwo is true iff n is a positive power of two.
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func DecimalStringWrong(number float64) string {
	return strconv.FormatFloat(number, 'G', -1, 64)
}

func BinaryStringCorrect(number float64) string {
	return PrintBinary(number)
}

func DoubleToHexBits(number float64) string {
	return fmt.Sprintf("%016x", math.Float64bits(number))
}

func DemonstrateMismatch(number float64) FloatDemo {
	return FloatDemo{
		DecimalStr:     DecimalStringWrong(number),
		BinaryFraction: BinaryStringCorrect(number),
		Ieee754Hex:     DoubleToHexBits(number),
	}
}

func RunQ505() {
	fmt.Printf("8 is power of two: %t\n", IsPowerOfTwo(8))
	fmt.Printf("6 is power of two: %t\n", IsPowerOfTwo(6))
	for _, number := range []float64{0.625, 0.1, 1.0 / 3.0} {
		demo := DemonstrateMismatch(number)
		fmt.Printf("n=%v decimal=%s binary=%s hex=%s\n", number, demo.DecimalStr, demo.BinaryFraction, demo.Ieee754Hex)
	}
}
