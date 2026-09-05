package chapter10

import "fmt"

func SearchRotated(values []int, target int) int {
	return searchRotated(values, 0, len(values)-1, target)
}

func searchRotated(values []int, left, right, target int) int {
	if right < left {
		return -1
	}
	mid := left + (right-left)/2
	if values[mid] == target {
		return mid
	}
	if values[left] < values[mid] {
		if target >= values[left] && target < values[mid] {
			return searchRotated(values, left, mid-1, target)
		}
		return searchRotated(values, mid+1, right, target)
	}
	if values[mid] < values[left] {
		if target > values[mid] && target <= values[right] {
			return searchRotated(values, mid+1, right, target)
		}
		return searchRotated(values, left, mid-1, target)
	}
	if values[left] != values[mid] {
		return -1
	}
	if values[mid] != values[right] {
		return searchRotated(values, mid+1, right, target)
	}
	leftResult := searchRotated(values, left, mid-1, target)
	if leftResult != -1 {
		return leftResult
	}
	return searchRotated(values, mid+1, right, target)
}

func RunQ1003() {
	values := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	for _, target := range []int{8, 1, 4, 5, 10} {
		fmt.Printf("%d -> %d\n", target, SearchRotated(values, target))
	}
}
