package chapter08

import "fmt"

func Multiply(a, b int, answer ...int) int {
	validateNonNegative(a, b)
	ans := 0
	if len(answer) > 0 {
		ans = answer[0]
	}
	if ans == 0 && a != 0 && b != 0 {
		ans = a
	}
	if a == 1 || b == 1 {
		return ans
	}
	if a == 0 || b == 0 {
		return 0
	}
	return Multiply(a, b-1, ans+a)
}

func MinProduct(a, b int) int {
	validateNonNegative(a, b)
	bigger, smaller := a, b
	if a < b {
		bigger, smaller = b, a
	}
	return minProductHelper(smaller, bigger)
}

func minProductHelper(smaller, bigger int) int {
	if smaller == 0 {
		return 0
	}
	if smaller == 1 {
		return bigger
	}
	half := smaller >> 1
	side1 := MinProduct(half, bigger)
	side2 := side1
	if smaller%2 == 1 {
		side2 = minProductHelper(smaller-half, bigger)
	}
	return side1 + side2
}

func MinProduct2(a, b int) int {
	validateNonNegative(a, b)
	bigger, smaller := a, b
	if a < b {
		bigger, smaller = b, a
	}
	return minProduct2Helper(smaller, bigger, map[int]int{})
}

func minProduct2Helper(smaller, bigger int, memo map[int]int) int {
	if smaller == 0 {
		return 0
	}
	if smaller == 1 {
		return bigger
	}
	if cached, ok := memo[smaller]; ok {
		return cached
	}
	half := smaller >> 1
	side1 := minProduct2Helper(half, bigger, memo)
	side2 := side1
	if smaller%2 == 1 {
		side2 = minProduct2Helper(smaller-half, bigger, memo)
	}
	memo[smaller] = side1 + side2
	return memo[smaller]
}

func MinProduct3(a, b int) int {
	validateNonNegative(a, b)
	bigger, smaller := a, b
	if a < b {
		bigger, smaller = b, a
	}
	return minProduct3Helper(smaller, bigger)
}

func minProduct3Helper(smaller, bigger int) int {
	if smaller == 0 {
		return 0
	}
	if smaller == 1 {
		return bigger
	}
	half := smaller >> 1
	halfProd := minProduct3Helper(half, bigger)
	if smaller%2 == 0 {
		return halfProd + halfProd
	}
	return halfProd + halfProd + bigger
}

func MultiplyBitBased(a, b int) int {
	validateNonNegative(a, b)
	product := 0
	shift := 0
	for bits := uint(b); bits != 0; bits, shift = bits>>1, shift+1 {
		if bits&1 != 0 {
			product += a << shift
		}
	}
	return product
}

func validateNonNegative(a, b int) {
	if a < 0 || b < 0 {
		panic("a and b must be non-negative")
	}
}

func RunQ805() {
	fmt.Printf("5*6 naive=%d min=%d memo=%d opt=%d bits=%d\n", Multiply(5, 6), MinProduct(5, 6), MinProduct2(5, 6), MinProduct3(5, 6), MultiplyBitBased(5, 6))
}
