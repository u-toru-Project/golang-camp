package chapter16

import "fmt"

func SubSort(arr []int) (int, int) {
	n := len(arr)
	if n <= 1 {
		return -1, -1
	}
	endLeft := 0
	leftMax := arr[0]
	for endLeft < n && arr[endLeft] >= leftMax {
		leftMax = arr[endLeft]
		endLeft++
	}
	if endLeft == n {
		return -1, -1
	}
	startRight := n - 1
	rightMin := arr[n-1]
	for startRight >= 0 && arr[startRight] <= rightMin {
		rightMin = arr[startRight]
		startRight--
	}
	if endLeft > startRight {
		endLeft = 0
		startRight = n - 1
	}
	midMin := arr[endLeft]
	midMax := arr[endLeft]
	for i := endLeft; i <= startRight; i++ {
		if arr[i] < midMin {
			midMin = arr[i]
		}
		if arr[i] > midMax {
			midMax = arr[i]
		}
	}
	for endLeft > 0 && arr[endLeft-1] > midMin {
		endLeft--
	}
	for startRight < n-1 && arr[startRight+1] < midMax {
		startRight++
	}
	return endLeft, startRight
}

func RunQ1616() {
	arr := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	start, end := SubSort(arr)
	fmt.Printf("Sort [%d, %d]\n", start, end)
}
