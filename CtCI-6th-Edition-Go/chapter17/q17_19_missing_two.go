package chapter17

import (
	"fmt"
	"math"
)

func MissingTwoSumSquares(arr []int, n int) (int, int) {
	validateMissingTwo(arr, n)
	expectedSum := int64(n) * int64(n+1) / 2
	expectedSq := int64(n) * int64(n+1) * (2*int64(n) + 1) / 6
	var actualSum, actualSq int64
	for _, value := range arr {
		actualSum += int64(value)
		actualSq += int64(value) * int64(value)
	}
	deltaSum := expectedSum - actualSum
	deltaSq := expectedSq - actualSq
	discriminant := 2*deltaSq - deltaSum*deltaSum
	if discriminant < 0 {
		panic("invalid input: no two missing numbers in 0..n")
	}
	b := int((deltaSum + integerSqrt(discriminant)) / 2)
	a := int(deltaSum - int64(b))
	if a < b {
		return a, b
	}
	return b, a
}

func MissingTwoXorBit(arr []int, n int) (int, int) {
	validateMissingTwo(arr, n)
	xorAll := 0
	for i := 0; i <= n; i++ {
		xorAll ^= i
	}
	for _, value := range arr {
		xorAll ^= value
	}
	if xorAll == 0 {
		panic("invalid input")
	}
	bit := xorAll & -xorAll
	missingA := 0
	for i := 0; i <= n; i++ {
		if i&bit != 0 {
			missingA ^= i
		}
	}
	for _, value := range arr {
		if value&bit != 0 {
			missingA ^= value
		}
	}
	missingB := xorAll ^ missingA
	if missingA < missingB {
		return missingA, missingB
	}
	return missingB, missingA
}

func validateMissingTwo(arr []int, n int) {
	if n < 2 {
		panic("n must be >= 2")
	}
	if len(arr) != n-1 {
		panic("expected array length mismatch")
	}
}

func integerSqrt(value int64) int64 {
	root := int64(math.Sqrt(float64(value)))
	for root*root > value {
		root--
	}
	for (root+1)*(root+1) <= value {
		root++
	}
	return root
}

func RunQ1719() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a, b := MissingTwoSumSquares(arr, 9)
	fmt.Println(a, b)
	a, b = MissingTwoXorBit(arr, 9)
	fmt.Println(a, b)
}
