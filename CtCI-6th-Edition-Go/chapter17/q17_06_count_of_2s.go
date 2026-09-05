package chapter17

import "fmt"

func CountDigitInRange(n, digit int) int {
	if n < 0 {
		panic("n must be >= 0")
	}
	if digit < 0 || digit > 9 {
		panic("digit must be 0..9")
	}
	power := 1
	total := 0
	for n/power > 0 {
		powerNext := power * 10
		higher := n / powerNext
		total += higher * power
		remainder := n % powerNext
		currentDigit := remainder / power
		if currentDigit > digit {
			total += power
		} else if currentDigit == digit {
			total += remainder%power + 1
		}
		power = powerNext
	}
	return total
}

func CountOfTwos(n int) int {
	return CountDigitInRange(n, 2)
}

func RunQ1706() {
	for _, n := range []int{0, 2, 20, 25, 100, 999} {
		fmt.Printf("twos through %d: %d\n", n, CountOfTwos(n))
	}
}
