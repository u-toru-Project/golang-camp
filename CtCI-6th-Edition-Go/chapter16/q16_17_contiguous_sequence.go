package chapter16

import "fmt"

func MaxContiguousSum(arr []int) int {
	if len(arr) == 0 {
		panic("Array must be non-empty.")
	}
	best := arr[0]
	current := 0
	for _, value := range arr {
		if value > current+value {
			current = value
		} else {
			current += value
		}
		if current > best {
			best = current
		}
	}
	return best
}

func RunQ1617() {
	fmt.Println(MaxContiguousSum([]int{2, -8, 3, -2, 4, -10}))
}
