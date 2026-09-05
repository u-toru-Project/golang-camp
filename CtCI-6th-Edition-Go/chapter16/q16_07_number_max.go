package chapter16

import "fmt"

func GetMax(a, b int) int {
	k := int((int64(a)-int64(b))>>63) & 1
	return a*(1-k) + b*k
}

func CountMaxes(values []int) int {
	if len(values) == 0 {
		return 0
	}
	currentMax := values[0]
	count := 1
	for i := 1; i < len(values); i++ {
		value := values[i]
		if value > currentMax {
			currentMax = value
			count = 1
		} else if value == currentMax {
			count++
		}
	}
	return count
}

func RunQ1607() {
	fmt.Printf("GetMax(10, 20) = %d\n", GetMax(10, 20))
	fmt.Printf("CountMaxes = %d\n", CountMaxes([]int{1, 2, 3, 2, 3, 3}))
}
