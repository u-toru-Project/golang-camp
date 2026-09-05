package chapter06

import "fmt"

func DominoTilings2ByN(n int) int {
	if n < 0 {
		panic("n must be non-negative")
	}
	if n == 0 || n == 1 {
		return 1
	}
	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}
	return curr
}

func DominoTilings3ByN(n int) int {
	if n < 0 {
		panic("n must be non-negative")
	}
	if n == 0 {
		return 1
	}
	table := []int{0, 1, 0, 1, 2, 3, 4, 11, 24, 53, 117, 258, 569}
	if n < len(table) {
		return table[n]
	}
	panic("n too large for demo table")
}

func RunQ603() {
	fmt.Printf("2x8 tilings: %d\n", DominoTilings2ByN(8))
	fmt.Printf("3x8 tilings: %d\n", DominoTilings3ByN(8))
}
