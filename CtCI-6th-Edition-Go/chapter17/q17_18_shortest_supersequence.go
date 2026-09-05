package chapter17

import (
	"fmt"
	"math"
)

type Window struct {
	Left  int
	Right int
}

func MinWindowString(bigArray, smallArray string) *Window {
	big := make([]byte, len(bigArray))
	copy(big, bigArray)
	small := make([]byte, len(smallArray))
	copy(small, smallArray)
	return MinWindowBytes(big, small)
}

func MinWindowBytes(bigArray, smallArray []byte) *Window {
	if len(smallArray) == 0 {
		panic("small_array must be non-empty")
	}
	frequencies := make(map[byte]int)
	for _, item := range smallArray {
		frequencies[item]++
	}
	missing := len(smallArray)
	minWinLen := math.MaxInt32
	minWinLeft, minWinRight := -1, -1
	left := 0
	for right := range bigArray {
		rightValue := bigArray[right]
		count := frequencies[rightValue]
		if count > 0 {
			missing--
		}
		frequencies[rightValue] = count - 1
		for left <= right && missing == 0 {
			if right-left+1 < minWinLen {
				minWinLen = right - left + 1
				minWinLeft = left
				minWinRight = right
			}
			leftValue := bigArray[left]
			if frequencies[leftValue] == 0 {
				missing++
				frequencies[leftValue] = 1
			} else {
				frequencies[leftValue]++
			}
			left++
		}
	}
	if minWinLen == math.MaxInt32 {
		return nil
	}
	return &Window{Left: minWinLeft, Right: minWinRight}
}

func MinWindowInts(bigArray, smallArray []int) *Window {
	if len(smallArray) == 0 {
		panic("small_array must be non-empty")
	}
	frequencies := make(map[int]int)
	for _, item := range smallArray {
		frequencies[item]++
	}
	missing := len(smallArray)
	minWinLen := math.MaxInt32
	minWinLeft, minWinRight := -1, -1
	left := 0
	for right := range bigArray {
		rightValue := bigArray[right]
		count := frequencies[rightValue]
		if count > 0 {
			missing--
		}
		frequencies[rightValue] = count - 1
		for left <= right && missing == 0 {
			if right-left+1 < minWinLen {
				minWinLen = right - left + 1
				minWinLeft = left
				minWinRight = right
			}
			leftValue := bigArray[left]
			if frequencies[leftValue] == 0 {
				missing++
				frequencies[leftValue] = 1
			} else {
				frequencies[leftValue]++
			}
			left++
		}
	}
	if minWinLen == math.MaxInt32 {
		return nil
	}
	return &Window{Left: minWinLeft, Right: minWinRight}
}

func RunQ1718() {
	window := MinWindowString("75902135791158897", "159")
	if window == nil {
		fmt.Println("(none)")
		return
	}
	fmt.Printf("%d, %d\n", window.Left, window.Right)
}
