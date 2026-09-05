package chapter17

import (
	"fmt"
	"math/rand"
)

func RandomNumberGenerator(lower, higher int) int {
	if lower > higher {
		panic("lower must be <= higher")
	}
	return lower + rand.Intn(higher-lower+1)
}

func ShuffleListRecursively(cards []int, currentIndex int) []int {
	if currentIndex < 0 || currentIndex >= len(cards) {
		panic("current_index out of range")
	}
	if currentIndex == 0 {
		return cards
	}
	ShuffleListRecursively(cards, currentIndex-1)
	randomIndex := RandomNumberGenerator(0, currentIndex)
	cards[randomIndex], cards[currentIndex] = cards[currentIndex], cards[randomIndex]
	return cards
}

func ShuffleListIteratively(cards []int) []int {
	for i := len(cards) - 1; i > 0; i-- {
		randomIndex := RandomNumberGenerator(0, i)
		cards[randomIndex], cards[i] = cards[i], cards[randomIndex]
	}
	return cards
}

func RunQ1702() {
	cards := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", cards)
	iter := append([]int(nil), cards...)
	fmt.Printf("Iterative: %v\n", ShuffleListIteratively(iter))
}
