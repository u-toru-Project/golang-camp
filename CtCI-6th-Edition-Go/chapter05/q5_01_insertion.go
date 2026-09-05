package chapter05

import (
	"fmt"
	"strconv"
)

// UpdateBits inserts m into n between bits i and j inclusive.
// Time O(1), space O(1).
func UpdateBits(n, m, i, j int) int {
	allOnes := int32(-1)
	var left int32
	if j == 31 {
		left = 0
	} else {
		left = allOnes << (j + 1)
	}
	right := int32(1<<i) - 1
	mask := left | right
	nCleared := int32(n) & mask
	mShifted := int32(m) << i
	return int(nCleared | mShifted)
}

func RunQ501() {
	n := 0b10000000000
	m := 0b10011
	fmt.Printf("%s, %s, i=2, j=6 -> %s\n", strconv.FormatInt(int64(n), 2), strconv.FormatInt(int64(m), 2), strconv.FormatInt(int64(UpdateBits(n, m, 2, 6)), 2))
}
