package chapter07

import "fmt"

const (
	EdgeFlat  = 0
	EdgeTab   = 1
	EdgeBlank = -1
)

type Piece struct {
	PieceId int
	Top     int
	Right   int
	Bottom  int
	Left    int
}

func NewPiece(pieceId, top, right, bottom, left int) *Piece {
	if pieceId < 0 {
		panic("invalid piece id")
	}
	return &Piece{PieceId: pieceId, Top: top, Right: right, Bottom: bottom, Left: left}
}

func (p *Piece) RotatedEdges(quarterTurns int) (top, right, bottom, left int) {
	turns := ((quarterTurns % 4) + 4) % 4
	top, right, bottom, left = p.Top, p.Right, p.Bottom, p.Left
	for range turns {
		top, right, bottom, left = left, top, right, bottom
	}
	return
}

func EdgesCompatible(a, b int) bool {
	return a+b == 0 && a != EdgeFlat
}

func CanPlace(placed *Piece, placedRow, placedCol int, candidate *Piece, candidateRow, candidateCol int, rotation ...int) bool {
	if placed == nil || candidate == nil {
		panic("piece is nil")
	}
	rot := 0
	if len(rotation) > 0 {
		rot = rotation[0]
	}
	dr := candidateRow - placedRow
	dc := candidateCol - placedCol
	if !((dr == 0 && dc == 1) || (dr == 0 && dc == -1) || (dr == 1 && dc == 0) || (dr == -1 && dc == 0)) {
		return false
	}
	ceTop, ceRight, ceBottom, ceLeft := candidate.RotatedEdges(rot)
	switch {
	case dr == 0 && dc == 1:
		return EdgesCompatible(placed.Right, ceLeft)
	case dr == 0 && dc == -1:
		return EdgesCompatible(placed.Left, ceRight)
	case dr == 1 && dc == 0:
		return EdgesCompatible(placed.Bottom, ceTop)
	default:
		return EdgesCompatible(placed.Top, ceBottom)
	}
}

func RunQ706() {
	left := NewPiece(1, EdgeFlat, EdgeTab, EdgeFlat, EdgeFlat)
	right := NewPiece(2, EdgeFlat, EdgeFlat, EdgeFlat, EdgeBlank)
	fmt.Printf("Can place: %t\n", CanPlace(left, 0, 0, right, 0, 1))
}
