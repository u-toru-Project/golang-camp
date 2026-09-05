package chapter08

import "fmt"

const NotFound = -1

func MagicIndex(array []int, bounds ...int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	if len(bounds) >= 1 {
		minIndex = bounds[0]
	}
	if len(bounds) >= 2 {
		maxIndex = bounds[1]
	}
	if maxIndex < minIndex {
		return NotFound
	}
	mid := (maxIndex + minIndex) / 2
	if array[mid] == mid {
		return mid
	}
	if array[mid] < mid {
		return MagicIndex(array, mid+1, maxIndex)
	}
	return MagicIndex(array, minIndex, mid-1)
}

func MagicIndexNonDistinct(array []int, bounds ...int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	if len(bounds) >= 1 {
		minIndex = bounds[0]
	}
	if len(bounds) >= 2 {
		maxIndex = bounds[1]
	}
	if maxIndex < minIndex {
		return NotFound
	}
	mid := (maxIndex + minIndex) / 2
	if array[mid] == mid {
		return mid
	}
	leftIndex := min(array[mid], mid-1)
	left := MagicIndexNonDistinct(array, minIndex, leftIndex)
	if left >= 0 {
		return left
	}
	rightIndex := max(array[mid], mid+1)
	return MagicIndexNonDistinct(array, rightIndex, maxIndex)
}

func RunQ803() {
	distinct := []int{-14, -12, 0, 1, 2, 5, 9, 10, 23, 25}
	duplicates := []int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}
	fmt.Printf("distinct=%d duplicates=%d\n", MagicIndex(distinct), MagicIndexNonDistinct(duplicates))
}
