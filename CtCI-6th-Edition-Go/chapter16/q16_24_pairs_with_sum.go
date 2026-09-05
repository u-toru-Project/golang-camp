package chapter16

import "fmt"

func CountPairsWithSum(arr []int, k int) int {
	left, right := 0, len(arr)-1
	count := 0
	for left < right {
		total := arr[left] + arr[right]
		if total < k {
			left++
		} else if total > k {
			right--
		} else if arr[left] == arr[right] {
			n := right - left + 1
			count += n * (n - 1) / 2
			break
		} else {
			leftValue := arr[left]
			rightValue := arr[right]
			leftRun, rightRun := 1, 1
			left++
			for left <= right && arr[left] == leftValue {
				leftRun++
				left++
			}
			right--
			for right >= left && arr[right] == rightValue {
				rightRun++
				right--
			}
			count += leftRun * rightRun
		}
	}
	return count
}

func RunQ1624() {
	fmt.Println(CountPairsWithSum([]int{4, 6, 10, 15, 16}, 21))
}
