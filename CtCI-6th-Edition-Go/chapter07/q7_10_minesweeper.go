package chapter07

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Hidden = -1
	Mine   = -2
)

type MinesweeperGame struct {
	Rows    int
	Cols    int
	rng     *rand.Rand
	mines   map[[2]int]struct{}
	visible [][]int
}

func NewMinesweeperGame(rows, cols, mineCount int, rng *rand.Rand) *MinesweeperGame {
	if rows < 1 || cols < 1 {
		panic("invalid dimensions")
	}
	if mineCount < 0 || mineCount >= rows*cols {
		panic("invalid mine count")
	}
	if rng == nil {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	visible := make([][]int, rows)
	for r := range rows {
		visible[r] = make([]int, cols)
		for c := range cols {
			visible[r][c] = Hidden
		}
	}
	g := &MinesweeperGame{Rows: rows, Cols: cols, rng: rng, mines: map[[2]int]struct{}{}, visible: visible}
	g.placeMines(mineCount)
	return g
}

func (g *MinesweeperGame) Reveal(row, col int) int {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols {
		panic("out of bounds")
	}
	if g.visible[row][col] != Hidden {
		panic("already revealed")
	}
	if _, ok := g.mines[[2]int{row, col}]; ok {
		g.visible[row][col] = Mine
		return Mine
	}
	return g.floodReveal(row, col)
}

func (g *MinesweeperGame) Cell(row, col int) int { return g.visible[row][col] }

func (g *MinesweeperGame) placeMines(count int) {
	cells := make([][2]int, 0, g.Rows*g.Cols)
	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			cells = append(cells, [2]int{r, c})
		}
	}
	for i := len(cells) - 1; i > 0; i-- {
		j := g.rng.Intn(i + 1)
		cells[i], cells[j] = cells[j], cells[i]
	}
	for i := range count {
		g.mines[cells[i]] = struct{}{}
	}
}

func (g *MinesweeperGame) countAdjacent(row, col int) int {
	total := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			r, c := row+dr, col+dc
			if r >= 0 && r < g.Rows && c >= 0 && c < g.Cols {
				if _, ok := g.mines[[2]int{r, c}]; ok {
					total++
				}
			}
		}
	}
	return total
}

func (g *MinesweeperGame) floodReveal(row, col int) int {
	stack := [][2]int{{row, col}}
	revealed := 0
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r, c := cur[0], cur[1]
		if g.visible[r][c] != Hidden {
			continue
		}
		if _, ok := g.mines[[2]int{r, c}]; ok {
			continue
		}
		count := g.countAdjacent(r, c)
		g.visible[r][c] = count
		revealed++
		if count == 0 {
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					nr, nc := r+dr, c+dc
					if nr >= 0 && nr < g.Rows && nc >= 0 && nc < g.Cols {
						stack = append(stack, [2]int{nr, nc})
					}
				}
			}
		}
	}
	return revealed
}

func RunQ710() {
	game := NewMinesweeperGame(4, 4, 3, rand.New(rand.NewSource(0)))
	fmt.Printf("Reveal (0,0): %d\n", game.Reveal(0, 0))
}
