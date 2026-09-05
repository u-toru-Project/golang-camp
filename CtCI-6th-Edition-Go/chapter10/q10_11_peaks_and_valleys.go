package chapter10

import (
	"fmt"
	"math"
	"sort"
)

func SortValleyPeak(values []int) {
	sort.Ints(values)
	for i := 1; i < len(values); i += 2 {
		values[i-1], values[i] = values[i], values[i-1]
	}
}

func SortValleyPeak2(values []int) {
	for i := 1; i < len(values); i += 2 {
		biggestIndex := maxIndex(values, i-1, i, i+1)
		if i != biggestIndex {
			values[i], values[biggestIndex] = values[biggestIndex], values[i]
		}
	}
}

func SortValleyPeak3(values []int) {
	for i := 1; i < len(values); i += 2 {
		if values[i-1] < values[i] {
			values[i-1], values[i] = values[i], values[i-1]
		}
		if i+1 < len(values) && values[i+1] < values[i] {
			values[i+1], values[i] = values[i], values[i+1]
		}
	}
}

func maxIndex(values []int, a, b, c int) int {
	length := len(values)
	aValue, bValue, cValue := math.MinInt, math.MinInt, math.MinInt
	if a >= 0 && a < length {
		aValue = values[a]
	}
	if b >= 0 && b < length {
		bValue = values[b]
	}
	if c >= 0 && c < length {
		cValue = values[c]
	}
	max := aValue
	if bValue > max {
		max = bValue
	}
	if cValue > max {
		max = cValue
	}
	if aValue == max {
		return a
	}
	if bValue == max {
		return b
	}
	return c
}

func RunQ1011() {
	samples := [][]int{
		{48, 40, 31, 62, 28, 21, 64, 40, 23, 17},
		{48, 40, 31, 62, 28, 21, 64, 40, 23, 17},
		{48, 40, 31, 62, 28, 21, 64, 40, 23, 17},
	}
	SortValleyPeak(samples[0])
	SortValleyPeak2(samples[1])
	SortValleyPeak3(samples[2])
	fmt.Printf("Sort:     %v\n", samples[0])
	fmt.Printf("Linear:   %v\n", samples[1])
	fmt.Printf("Adjacent: %v\n", samples[2])
}
