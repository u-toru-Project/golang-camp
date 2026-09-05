package chapter16

import "fmt"

func Negate(n int) int {
	if n == 0 {
		return 0
	}
	step := -1
	if n < 0 {
		step = 1
	}
	result := 0
	for n != 0 {
		n += step
		result += step
	}
	return result
}

func Subtract(a, b int) int {
	return a + Negate(b)
}

func Multiply(a, b int) int {
	negative := b < 0
	absoluteB := b
	if negative {
		absoluteB = Negate(b)
	}
	total := 0
	for i := 0; i < absoluteB; i++ {
		total += a
	}
	if negative {
		return Negate(total)
	}
	return total
}

func Divide(a, b int) int {
	if b == 0 {
		panic("Division by zero.")
	}
	negativeResult := (a < 0) != (b < 0)
	absoluteA := a
	if a < 0 {
		absoluteA = Negate(a)
	}
	absoluteB := b
	if b < 0 {
		absoluteB = Negate(b)
	}
	quotient := 0
	accumulated := 0
	for accumulated+absoluteB <= absoluteA {
		accumulated += absoluteB
		quotient++
	}
	if negativeResult {
		if accumulated != absoluteA {
			quotient++
		}
		quotient = Negate(quotient)
	}
	return quotient
}

func RunQ1609() {
	fmt.Printf("Subtract(10, 3) = %d\n", Subtract(10, 3))
	fmt.Printf("Multiply(10, -3) = %d\n", Multiply(10, -3))
	fmt.Printf("Divide(-10, 3) = %d\n", Divide(-10, 3))
}
