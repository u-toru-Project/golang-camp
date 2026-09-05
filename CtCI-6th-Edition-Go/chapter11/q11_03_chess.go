package chapter11

import (
	"fmt"
	"strings"
	"unicode"
)

const BoardSize = 8

func InBounds(row, col int) bool {
	return row >= 0 && row < BoardSize && col >= 0 && col < BoardSize
}

func ParseBoard(fenPieces string) [][]*string {
	rows := strings.Split(strings.TrimSpace(fenPieces), "/")
	if len(rows) != BoardSize {
		panic("expected 8 ranks")
	}

	board := make([][]*string, BoardSize)
	for r, rank := range rows {
		row := make([]*string, 0, BoardSize)
		for _, ch := range rank {
			if unicode.IsDigit(ch) {
				for i := 0; i < int(ch-'0'); i++ {
					row = append(row, nil)
				}
			} else {
				piece := string(ch)
				row = append(row, &piece)
			}
		}
		if len(row) != BoardSize {
			panic("invalid rank width")
		}
		board[r] = row
	}
	return board
}

func KingPosition(board [][]*string, king string) (int, int) {
	for row := range board {
		for col := range board[row] {
			if board[row][col] != nil && *board[row][col] == king {
				return row, col
			}
		}
	}
	panic("king not found")
}

func IsSquareAttacked(board [][]*string, row, col int, byWhite bool) bool {
	pawn := "p"
	direction := 1
	if byWhite {
		pawn = "P"
		direction = -1
	}
	for _, dc := range []int{-1, 1} {
		r := row + direction
		c := col + dc
		if InBounds(r, c) && board[r][c] != nil && *board[r][c] == pawn {
			return true
		}
	}

	rook := "r"
	if byWhite {
		rook = "R"
	}
	deltas := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range deltas {
		r := row + d[0]
		c := col + d[1]
		for InBounds(r, c) {
			piece := board[r][c]
			if piece != nil {
				if *piece == rook {
					return true
				}
				break
			}
			r += d[0]
			c += d[1]
		}
	}
	return false
}

func RunQ1103() {
	board := ParseBoard("8/8/8/8/8/8/8/k7")
	kr, kc := KingPosition(board, "k")
	fmt.Printf("king at (%d,%d)\n", kr, kc)

	attacked := ParseBoard("8/8/8/8/8/8/r7/k7")
	fmt.Printf("rook attacks king: %t\n", IsSquareAttacked(attacked, 7, 0, false))
}
