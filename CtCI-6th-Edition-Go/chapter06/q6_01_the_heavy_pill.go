package chapter06

import (
	"fmt"
	"math"
)

func MinWeighingsBalance(numPills int) int {
	if numPills < 1 {
		panic("num_pills must be positive")
	}
	if numPills == 1 {
		return 0
	}
	return int(math.Ceil(math.Log(float64(numPills)) / math.Log(3)))
}

func FindHeavyPillIndex(pills []float64) int {
	if pills == nil {
		panic("pills is nil")
	}
	if len(pills) < 1 {
		panic("need at least one pill")
	}
	heavyCount := 0
	for _, p := range pills {
		if p > 1.0 {
			heavyCount++
		}
	}
	if heavyCount != 1 {
		panic("expected exactly one heavy pill")
	}
	return findHeavyRecursive(pills, 0, len(pills)-1)
}

func findHeavyRecursive(pills []float64, low, high int) int {
	if low == high {
		return low
	}
	n := high - low + 1
	group := (n + 2) / 3
	leftEnd := low + group - 1
	midEnd := leftEnd + group
	leftWeight := sumRange(pills, low, leftEnd)
	midWeight := 0.0
	if midEnd >= leftEnd+1 {
		midWeight = sumRange(pills, leftEnd+1, midEnd)
	}
	if leftWeight > float64(group) {
		return findHeavyRecursive(pills, low, leftEnd)
	}
	if midWeight > float64(group) {
		return findHeavyRecursive(pills, leftEnd+1, midEnd)
	}
	return findHeavyRecursive(pills, midEnd+1, high)
}

func sumRange(pills []float64, start, end int) float64 {
	sum := 0.0
	last := min(end, len(pills)-1)
	for i := start; i <= last; i++ {
		sum += pills[i]
	}
	return sum
}

func RunQ601() {
	fmt.Printf("8 pills need %d weighings\n", MinWeighingsBalance(8))
	pills := []float64{1.0, 1.0, 1.0, 1.0, 1.1, 1.0, 1.0, 1.0}
	fmt.Printf("Heavy pill at index %d\n", FindHeavyPillIndex(pills))
}
