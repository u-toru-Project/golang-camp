package chapter16

import "fmt"

type Piece int

const (
	PieceNA Piece = iota
	PieceX
	PieceO
)

const BoardSize3 = 3

type TicTacToe struct {
	board         [BoardSize3][BoardSize3]Piece
	winner        Piece
	currentPlayer Piece
}

func NewTicTacToe(currentPlayer Piece) *TicTacToe {
	if currentPlayer == PieceNA {
		panic("The starting player must be X or O.")
	}
	return &TicTacToe{currentPlayer: currentPlayer}
}

func (g *TicTacToe) Place(row, column int, piece Piece) {
	if row < 0 || row >= BoardSize3 || column < 0 || column >= BoardSize3 {
		panic("Location is out of bounds.")
	}
	if piece != g.currentPlayer {
		panic("Wrong player.")
	}
	if g.board[row][column] != PieceNA {
		panic("Square already occupied.")
	}
	if g.winner != PieceNA {
		panic("Game has already ended.")
	}
	g.board[row][column] = piece
	if piece == PieceX {
		g.currentPlayer = PieceO
	} else {
		g.currentPlayer = PieceX
	}
	g.winner = g.winnerFromMove(row, column)
}

func (g *TicTacToe) GetWinner() Piece {
	return g.winner
}

func (g *TicTacToe) winnerFromMove(row, column int) Piece {
	piece := g.board[row][column]
	if g.hasWonRow(row) || g.hasWonColumn(column) || g.hasWonDiagonal() || g.hasWonAntiDiagonal() {
		return piece
	}
	return PieceNA
}

func (g *TicTacToe) hasWonRow(row int) bool {
	return isLine(g.board[row][0], g.board[row][1], g.board[row][2])
}

func (g *TicTacToe) hasWonColumn(column int) bool {
	return isLine(g.board[0][column], g.board[1][column], g.board[2][column])
}

func (g *TicTacToe) hasWonDiagonal() bool {
	return isLine(g.board[0][0], g.board[1][1], g.board[2][2])
}

func (g *TicTacToe) hasWonAntiDiagonal() bool {
	return isLine(g.board[0][2], g.board[1][1], g.board[2][0])
}

func HasWon(board [][]Piece) Piece {
	for i := range BoardSize3 {
		if isLine(board[i][0], board[i][1], board[i][2]) {
			return board[i][0]
		}
		if isLine(board[0][i], board[1][i], board[2][i]) {
			return board[0][i]
		}
	}
	if isLine(board[0][0], board[1][1], board[2][2]) {
		return board[0][0]
	}
	if isLine(board[0][2], board[1][1], board[2][0]) {
		return board[0][2]
	}
	return PieceNA
}

func isLine(first, second, third Piece) bool {
	return first != PieceNA && first == second && second == third
}

func RunQ1604() {
	game := NewTicTacToe(PieceX)
	game.Place(1, 1, PieceX)
	fmt.Printf("Winner after (1,1) X: %v\n", game.GetWinner())
}
