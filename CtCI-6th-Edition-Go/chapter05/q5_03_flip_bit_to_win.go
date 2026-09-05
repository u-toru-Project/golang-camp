package chapter05

import (
	"fmt"
	"strconv"
)

// FlipBitToWin returns the longest 1-run after flipping at most one 0.
// Time O(b), space O(b).
func FlipBitToWin(number int) (int, error) {
	if number < 0 {
		return 0, fmt.Errorf("number must be non-negative")
	}
	bits := strconv.FormatInt(int64(number), 2)
	maxCount := 0
	count := 0
	zeroSeen := false
	i := len(bits)
	lastZero := i
	for i > 0 {
		if bits[i-1] == '1' {
			count++
		} else if !zeroSeen {
			lastZero = i
			zeroSeen = true
		} else {
			if count > maxCount {
				maxCount = count
			}
			i = lastZero
			zeroSeen = false
			count = 0
		}
		i--
	}
	if count > maxCount {
		maxCount = count
	}
	return maxCount + 1, nil
}

// FlipBitToWinAlt tracks previous and current 1-runs around a single 0.
// Time O(b), space O(1).
func FlipBitToWinAlt(number int) (int, error) {
	if number < 0 {
		return 0, fmt.Errorf("number must be non-negative")
	}
	longest := 1
	currentSegment := 0
	pastSegment := 0
	bits := uint32(number)
	for bits != 0 {
		if bits&1 != 0 {
			currentSegment++
		} else {
			if bits&2 == 0 {
				pastSegment = 0
			} else {
				pastSegment = currentSegment
			}
			currentSegment = 0
		}
		if currentSegment+pastSegment+1 > longest {
			longest = currentSegment + pastSegment + 1
		}
		bits >>= 1
	}
	return longest, nil
}

func RunQ503() {
	for _, number := range []int{0b0, 0b111, 0b10011100111, 0b10110110111, 0b11011101111} {
		a, _ := FlipBitToWin(number)
		b, _ := FlipBitToWinAlt(number)
		fmt.Printf("%s: %d / %d\n", strconv.FormatInt(int64(number), 2), a, b)
	}
}
