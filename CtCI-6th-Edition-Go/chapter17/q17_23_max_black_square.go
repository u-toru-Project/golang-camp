package chapter17

import "fmt"

type SquareCorner struct {
	Row, Column int
}

type BlackSquare struct {
	TopLeft     SquareCorner
	BottomRight SquareCorner
}

func MaxBlackSquare(matrix [][]int) *BlackSquare {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	rows := len(matrix)
	cols := len(matrix[0])
	for _, row := range matrix {
		if len(row) != cols {
			panic("matrix must be rectangular")
		}
	}
	down, right := precomputeExtensions(matrix, rows, cols)
	maxSide := min(cols, rows)
	for side := maxSide; side >= 1; side-- {
		for top := 0; top <= rows-side; top++ {
			for left := 0; left <= cols-side; left++ {
				if isValidSquare(top, left, side, down, right) {
					return &BlackSquare{
						TopLeft:     SquareCorner{top, left},
						BottomRight: SquareCorner{top + side - 1, left + side - 1},
					}
				}
			}
		}
	}
	return nil
}

func precomputeExtensions(matrix [][]int, rows, cols int) ([][]int, [][]int) {
	down := make([][]int, rows+1)
	right := make([][]int, rows+1)
	for row := 0; row <= rows; row++ {
		down[row] = make([]int, cols+1)
		right[row] = make([]int, cols+1)
	}
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if matrix[row][col] == 0 {
				continue
			}
			down[row][col] = 1 + down[row+1][col]
			right[row][col] = 1 + right[row][col+1]
		}
	}
	return down, right
}

func isValidSquare(topRow, leftCol, side int, down, right [][]int) bool {
	bottomRow := topRow + side - 1
	rightCol := leftCol + side - 1
	if down[topRow][leftCol] < side || right[topRow][leftCol] < side {
		return false
	}
	return down[topRow][rightCol] >= side && right[bottomRow][leftCol] >= side
}

func RunQ1723() {
	matrix := [][]int{
		{0, 1, 1, 1, 1},
		{1, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 1},
	}
	fmt.Println(MaxBlackSquare(matrix))
}
