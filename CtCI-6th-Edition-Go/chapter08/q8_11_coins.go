package chapter08

import "fmt"

var StandardCoinSizes = []int{1, 5, 10, 25}

func CoinCombinations(amount int, coinSizes ...[]int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 {
		return 0
	}
	sizes := StandardCoinSizes
	if len(coinSizes) > 0 {
		sizes = coinSizes[0]
	}
	if len(sizes) == 0 {
		return 0
	}
	last := sizes[len(sizes)-1]
	withoutLast := sizes[:len(sizes)-1]
	return CoinCombinations(amount, withoutLast) + CoinCombinations(amount-last, sizes)
}

func CoinCombinationsIterative(amount int, coinSizes ...[]int) int {
	if amount < 0 {
		panic("amount must be non-negative")
	}
	sizes := StandardCoinSizes
	if len(coinSizes) > 0 {
		sizes = coinSizes[0]
	}
	for _, coin := range sizes {
		if coin <= 0 {
			panic("coin sizes must be positive")
		}
	}
	table := make([]int, amount+1)
	table[0] = 1
	for _, coin := range sizes {
		for j := coin; j <= amount; j++ {
			table[j] += table[j-coin]
		}
	}
	return table[amount]
}

func RunQ811() {
	for _, amount := range []int{0, 1, 5, 10, 100} {
		rec := -1
		if amount <= 10 {
			rec = CoinCombinations(amount)
		}
		fmt.Printf("%d: recursive=%d dp=%d\n", amount, rec, CoinCombinationsIterative(amount))
	}
}
