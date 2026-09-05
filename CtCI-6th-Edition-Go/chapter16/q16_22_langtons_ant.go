package chapter16

import "fmt"

type Cell struct {
	Row, Col int
}

type BoundingBoxResult struct {
	MinRow, MaxRow, MinCol, MaxCol int
}

var antDirections = [][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

func Simulate(k int) map[Cell]bool {
	if k < 0 {
		panic("k must be non-negative.")
	}
	grid := make(map[Cell]bool)
	row, col, direction := 0, 0, 0
	for range k {
		black := grid[Cell{row, col}]
		if black {
			grid[Cell{row, col}] = false
			direction = (direction + 3) % 4
		} else {
			grid[Cell{row, col}] = true
			direction = (direction + 1) % 4
		}
		row += antDirections[direction][0]
		col += antDirections[direction][1]
	}
	return grid
}

func BoundingBox(grid map[Cell]bool) BoundingBoxResult {
	if len(grid) == 0 {
		return BoundingBoxResult{}
	}
	minRow, maxRow := 1<<31-1, -1<<31
	minCol, maxCol := 1<<31-1, -1<<31
	for cell := range grid {
		if cell.Row < minRow {
			minRow = cell.Row
		}
		if cell.Row > maxRow {
			maxRow = cell.Row
		}
		if cell.Col < minCol {
			minCol = cell.Col
		}
		if cell.Col > maxCol {
			maxCol = cell.Col
		}
	}
	return BoundingBoxResult{minRow, maxRow, minCol, maxCol}
}

func RunQ1622() {
	grid := Simulate(12)
	box := BoundingBox(grid)
	fmt.Printf("Cells=%d box=(%d,%d,%d,%d)\n", len(grid), box.MinRow, box.MaxRow, box.MinCol, box.MaxCol)
}
