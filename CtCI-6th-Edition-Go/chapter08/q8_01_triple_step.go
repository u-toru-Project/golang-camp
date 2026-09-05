package chapter08

import "fmt"

func TripleHop(stairs int) int {
	if stairs < 0 {
		return 0
	}
	if stairs == 0 || stairs == 1 {
		return 1
	}
	return TripleHop(stairs-1) + TripleHop(stairs-2) + TripleHop(stairs-3)
}

func TripleHopDp(stairs int) int {
	if stairs < 0 {
		return 0
	}
	memo := make([]int, stairs+1)
	for i := range memo {
		memo[i] = -1
	}
	return tripleHopFill(stairs, memo)
}

func tripleHopFill(stairs int, memo []int) int {
	if stairs < 0 {
		return 0
	}
	memo[0] = 1
	if stairs >= 1 {
		memo[1] = 1
	}
	if stairs >= 2 {
		memo[2] = memo[1] + memo[0]
	}
	if stairs > 2 {
		for i := 3; i <= stairs; i++ {
			memo[i] = memo[i-1] + memo[i-2] + memo[i-3]
		}
	}
	return memo[stairs]
}

func RunQ801() {
	for _, stairs := range []int{0, 1, 2, 3, 4, 5, 6} {
		fmt.Printf("%d: recursive=%d dp=%d\n", stairs, TripleHop(stairs), TripleHopDp(stairs))
	}
}
