package chapter01

import "fmt"

func IsSquareMatrix(matrix [][]int) bool {
	if matrix == nil {
		return false
	}
	size := len(matrix)
	for _, row := range matrix {
		if row == nil || len(row) != size {
			return false
		}
	}
	return true
}

// Rotate rotates 90 degrees clockwise in place.
// Time O(n^2), space O(1).
func Rotate(matrix [][]int) bool {
	if !IsSquareMatrix(matrix) {
		return false
	}
	size := len(matrix)
	for layer := 0; layer < size/2; layer++ {
		first, last := layer, size-1-layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-offset][first]
			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}
	return true
}

func RunQ107() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate(matrix)
	fmt.Println(matrix)
}
