package chapter17

import "fmt"

func MajorityCandidate(arr []int) int {
	if len(arr) == 0 {
		panic("arr must be non-empty")
	}
	candidate := arr[0]
	count := 0
	for _, value := range arr {
		if count == 0 {
			candidate = value
			count = 1
		} else if value == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func IsMajority(value int, arr []int) bool {
	count := 0
	for _, item := range arr {
		if item == value {
			count++
		}
	}
	return count > len(arr)/2
}

func MajorityElement(arr []int) int {
	candidate := MajorityCandidate(arr)
	if IsMajority(candidate, arr) {
		return candidate
	}
	return -1
}

func RunQ1710() {
	for _, sample := range [][]int{{4, 4, 4, 4, 5, 5, 5, 5, 5}, {1, 2, 3, 4, 5}, {7, 7, 7, 7, 1, 2, 3}} {
		fmt.Printf("%v -> %d\n", sample, MajorityElement(sample))
	}
}
