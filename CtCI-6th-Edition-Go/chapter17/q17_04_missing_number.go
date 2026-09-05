package chapter17

import "fmt"

func MissingNumberXor(arr []int, n int) int {
	validateMissing(arr, n)
	for _, v := range arr {
		if v < 0 || v > n {
			panic("elements must lie in 0..n")
		}
	}
	result := n
	for i, v := range arr {
		result ^= i ^ v
	}
	return result
}

func MissingNumberSum(arr []int, n int) int {
	validateMissing(arr, n)
	expected := n * (n + 1) / 2
	actual := 0
	for _, v := range arr {
		actual += v
	}
	return expected - actual
}

func validateMissing(arr []int, n int) {
	if n < 0 {
		panic("n must be >= 0")
	}
	if len(arr) != n {
		panic("expected array length mismatch")
	}
}

func RunQ1704() {
	arr := []int{0, 1, 2, 3, 5}
	fmt.Printf("xor=%d sum=%d\n", MissingNumberXor(arr, 5), MissingNumberSum(arr, 5))
}
