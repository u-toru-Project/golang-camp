package chapter07

import "fmt"

const (
	BoardSize    = 8
	OthelloEmpty = byte('.')
	OthelloBlack = byte('B')
	OthelloWhite = byte('W')
)

var othelloDirections = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

type OthelloGame struct {
	Board [][]byte
}

func NewOthelloGame() *OthelloGame {
	board := make([][]byte, BoardSize)
	for r := range BoardSize {
		board[r] = make([]byte, BoardSize)
		for c := range BoardSize {
			board[r][c] = OthelloEmpty
		}
	}
	mid := BoardSize / 2
	board[mid-1][mid-1] = OthelloWhite
	board[mid-1][mid] = OthelloBlack
	board[mid][mid-1] = OthelloBlack
	board[mid][mid] = OthelloWhite
	return &OthelloGame{Board: board}
}

func (g *OthelloGame) InBounds(row, col int) bool {
	return row >= 0 && row < BoardSize && col >= 0 && col < BoardSize
}

func (g *OthelloGame) FlipsInDirection(row, col int, player byte, dr, dc int) [][2]int {
	opponent := othelloOpponent(player)
	r, c := row+dr, col+dc
	captured := make([][2]int, 0)
	for g.InBounds(r, c) && g.Board[r][c] == opponent {
		captured = append(captured, [2]int{r, c})
		r += dr
		c += dc
	}
	if len(captured) > 0 && g.InBounds(r, c) && g.Board[r][c] == player {
		return captured
	}
	return nil
}

func (g *OthelloGame) ValidMove(row, col int, player byte) bool {
	if !g.InBounds(row, col) || g.Board[row][col] != OthelloEmpty {
		return false
	}
	for _, d := range othelloDirections {
		if len(g.FlipsInDirection(row, col, player, d[0], d[1])) > 0 {
			return true
		}
	}
	return false
}

func (g *OthelloGame) ApplyMove(row, col int, player byte) int {
	if player != OthelloBlack && player != OthelloWhite {
		panic("invalid player")
	}
	if !g.ValidMove(row, col, player) {
		panic("invalid move")
	}
	flipped := 0
	g.Board[row][col] = player
	for _, d := range othelloDirections {
		for _, pos := range g.FlipsInDirection(row, col, player, d[0], d[1]) {
			g.Board[pos[0]][pos[1]] = player
			flipped++
		}
	}
	return flipped
}

func othelloOpponent(player byte) byte {
	if player == OthelloBlack {
		return OthelloWhite
	}
	return OthelloBlack
}

func RunQ708() {
	game := NewOthelloGame()
	fmt.Printf("Opening move flips: %d\n", game.ApplyMove(2, 3, OthelloBlack))
}
