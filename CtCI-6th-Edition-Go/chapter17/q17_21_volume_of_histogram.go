package chapter17

import "fmt"

func FindVolume(histogram []int) int {
	for _, height := range histogram {
		if height < 0 {
			panic("bar heights must be non-negative")
		}
	}
	if len(histogram) <= 2 {
		return 0
	}
	n := len(histogram)
	heightLeft := make([]int, n)
	heightRight := make([]int, n)
	for i := 1; i < n; i++ {
		heightLeft[i] = max(heightLeft[i-1], histogram[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		heightRight[i] = max(heightRight[i+1], histogram[i+1])
	}
	volume := 0
	for i := 1; i < n-1; i++ {
		minHeight := min(heightRight[i], heightLeft[i])
		if minHeight-histogram[i] > 0 {
			volume += minHeight - histogram[i]
		}
	}
	return volume
}

func RunQ1721() {
	fmt.Println(FindVolume([]int{0, 0, 4, 0, 0, 6, 0, 0, 3, 0, 5, 0, 1, 0, 0, 0}))
}
