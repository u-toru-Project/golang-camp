package chapter16

import "fmt"

func DivingBoardLengths(k, longer, shorter int) []int {
	if k < 0 {
		panic("k must be non-negative.")
	}
	if shorter > longer {
		panic("Shorter must not exceed longer.")
	}
	if k == 0 {
		return []int{0}
	}
	if shorter == longer {
		return []int{longer * k}
	}
	lengths := make([]int, 0, k+1)
	for longCount := 0; longCount <= k; longCount++ {
		shortCount := k - longCount
		lengths = append(lengths, longCount*longer+shortCount*shorter)
	}
	return lengths
}

func RunQ1611() {
	fmt.Println(DivingBoardLengths(10, 10, 5))
}
