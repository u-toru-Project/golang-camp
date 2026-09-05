package chapter01

import "slices"

import "fmt"

func isEmpty(matrix [][]int) bool {
	return matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0
}

func nullifyRow(matrix [][]int, row int) {
	for column := range matrix[row] {
		matrix[row][column] = 0
	}
}

func nullifyColumn(matrix [][]int, column int) {
	for row := range matrix {
		matrix[row][column] = 0
	}
}

// SetZeros zeroes rows and columns that contain a zero.
// Time O(mn), space O(m+n).
func SetZeros(matrix [][]int) {
	if isEmpty(matrix) {
		return
	}
	rowHasZero := make([]bool, len(matrix))
	columnHasZero := make([]bool, len(matrix[0]))
	for row := range matrix {
		for column := range matrix[0] {
			if matrix[row][column] == 0 {
				rowHasZero[row] = true
				columnHasZero[column] = true
			}
		}
	}
	for row, has := range rowHasZero {
		if has {
			nullifyRow(matrix, row)
		}
	}
	for column, has := range columnHasZero {
		if has {
			nullifyColumn(matrix, column)
		}
	}
}

// SetZeros2 uses the first row and column as markers.
// Time O(mn), space O(1).
func SetZeros2(matrix [][]int) {
	if isEmpty(matrix) {
		return
	}
	firstRowHasZero := slices.Contains(matrix[0], 0)
	firstColumnHasZero := false
	for _, row := range matrix {
		if row[0] == 0 {
			firstColumnHasZero = true
			break
		}
	}
	for row := 1; row < len(matrix); row++ {
		for column := 1; column < len(matrix[0]); column++ {
			if matrix[row][column] == 0 {
				matrix[row][0] = 0
				matrix[0][column] = 0
			}
		}
	}
	for row := 1; row < len(matrix); row++ {
		if matrix[row][0] == 0 {
			nullifyRow(matrix, row)
		}
	}
	for column := 1; column < len(matrix[0]); column++ {
		if matrix[0][column] == 0 {
			nullifyColumn(matrix, column)
		}
	}
	if firstRowHasZero {
		nullifyRow(matrix, 0)
	}
	if firstColumnHasZero {
		nullifyColumn(matrix, 0)
	}
}

func RunQ108() {
	matrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 0, 9}}
	SetZeros(matrix)
	fmt.Println(matrix)
}
