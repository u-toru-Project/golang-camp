package chapter17

import (
	"fmt"
	"math"
)

func MaxSubmatrixSum(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		panic("matrix must be non-empty")
	}
	rows := len(matrix)
	cols := len(matrix[0])
	for _, row := range matrix {
		if len(row) != cols {
			panic("matrix must be rectangular")
		}
	}
	best := math.MinInt32
	for top := range rows {
		colSums := make([]int, cols)
		for bottom := top; bottom < rows; bottom++ {
			for col := range cols {
				colSums[col] += matrix[bottom][col]
			}
			sum := maxSubarraySum(colSums)
			if sum > best {
				best = sum
			}
		}
	}
	return best
}

func maxSubarraySum(arr []int) int {
	best := arr[0]
	current := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > current+arr[i] {
			current = arr[i]
		} else {
			current += arr[i]
		}
		if current > best {
			best = current
		}
	}
	return best
}

func RunQ1724() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, -5000, 5, 6},
		{10, 20, 30, 40, 50},
	}
	fmt.Println(MaxSubmatrixSum(matrix))
}
