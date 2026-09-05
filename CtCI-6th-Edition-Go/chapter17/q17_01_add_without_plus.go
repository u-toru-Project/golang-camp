package chapter17

import "fmt"

func AddWithoutPlus(a, b int) int {
	ensureNonNegative(a, b)
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func AddWithoutPlusRecursive(a, b int) int {
	ensureNonNegative(a, b)
	if b == 0 {
		return a
	}
	return AddWithoutPlusRecursive(a^b, (a&b)<<1)
}

func ensureNonNegative(a, b int) {
	if a < 0 || b < 0 {
		panic("this demo supports non-negative ints only")
	}
}

func RunQ1701() {
	for _, pair := range [][2]int{{1, 1}, {1, 2}, {1001, 234}, {0, 5}} {
		fmt.Printf("%d + %d = %d / %d\n", pair[0], pair[1], AddWithoutPlus(pair[0], pair[1]), AddWithoutPlusRecursive(pair[0], pair[1]))
	}
}
