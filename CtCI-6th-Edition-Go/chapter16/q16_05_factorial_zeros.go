package chapter16

import "fmt"

func TrailingZerosInFactorial(n int) int {
	if n < 0 {
		panic("n must be non-negative.")
	}
	count := 0
	for powerOfFive := int64(5); powerOfFive <= int64(n); powerOfFive *= 5 {
		count += int(int64(n) / powerOfFive)
	}
	return count
}

func RunQ1605() {
	for _, n := range []int{0, 5, 10, 25, 100} {
		fmt.Printf("%d! has %d trailing zeros\n", n, TrailingZerosInFactorial(n))
	}
}
