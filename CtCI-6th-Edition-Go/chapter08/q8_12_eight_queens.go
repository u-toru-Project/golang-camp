package chapter08

import "fmt"

func PositionIsValid(row, col int, board [][]rune) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		panic("board must be non-empty")
	}
	n := len(board)
	if row < 0 || row >= n || col < 0 || col >= len(board[row]) {
		panic("coordinates out of bounds")
	}
	for curRow := range row {
		if board[curRow][col] == 'Q' {
			return false
		}
	}
	for curCol := 0; curCol < len(board[row]); curCol++ {
		if board[row][curCol] == 'Q' {
			return false
		}
	}
	for curRow, curCol := row-1, col-1; curRow >= 0 && curCol >= 0; curRow, curCol = curRow-1, curCol-1 {
		if board[curRow][curCol] == 'Q' {
			return false
		}
	}
	for curRow, curCol := row-1, col+1; curRow >= 0 && curCol < len(board[row]); curRow, curCol = curRow-1, curCol+1 {
		if board[curRow][curCol] == 'Q' {
			return false
		}
	}
	return true
}

func Queens(n int) [][][]rune {
	if n < 1 {
		panic("n must be positive")
	}
	if n > 12 {
		panic("n exceeds safety bound")
	}
	arrangements := make([][][]rune, 0)
	board := make([][]rune, n)
	for i := range n {
		board[i] = make([]rune, n)
		for j := range n {
			board[i][j] = '.'
		}
	}
	var helper func(row int)
	helper = func(row int) {
		if row == n {
			clone := make([][]rune, n)
			for i := range n {
				clone[i] = append([]rune{}, board[i]...)
			}
			arrangements = append(arrangements, clone)
			return
		}
		for col := range n {
			if !PositionIsValid(row, col, board) {
				continue
			}
			board[row][col] = 'Q'
			helper(row + 1)
			board[row][col] = '.'
		}
	}
	helper(0)
	return arrangements
}

func RunQ812() {
	fmt.Printf("n=4 solutions=%d\n", len(Queens(4)))
	fmt.Printf("n=8 solutions=%d\n", len(Queens(8)))
}
