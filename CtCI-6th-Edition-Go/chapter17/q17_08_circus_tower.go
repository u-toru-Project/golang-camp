package chapter17

import (
	"fmt"
	"sort"
)

type HeightWeight struct {
	Height int
	Weight int
}

func FindMaxPeople(pairs []HeightWeight) int {
	if len(pairs) == 0 {
		return 0
	}
	sorted := append([]HeightWeight(nil), pairs...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Weight > sorted[j].Weight
	})
	dp := make([]int, len(sorted))
	for i := range dp {
		dp[i] = 1
	}
	maxSoFar := 1
	for i := 1; i < len(sorted); i++ {
		best := 1
		for j := 0; j < i; j++ {
			if sorted[i].Height < sorted[j].Height && 1+dp[j] > best {
				best = 1 + dp[j]
			}
		}
		dp[i] = best
		if dp[i] > maxSoFar {
			maxSoFar = dp[i]
		}
	}
	return maxSoFar
}

func RunQ1708() {
	people := []HeightWeight{{65, 100}, {70, 150}, {56, 90}, {75, 190}, {60, 95}, {68, 110}}
	fmt.Printf("Max tower: %d\n", FindMaxPeople(people))
}
