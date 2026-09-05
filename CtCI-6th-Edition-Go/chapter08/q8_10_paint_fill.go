package chapter08

import "fmt"

func FloodFill(screen [][]int, i, j, color, newColor int) {
	if i < 0 || i >= len(screen) || j < 0 || j >= len(screen[0]) || screen[i][j] != color {
		return
	}
	screen[i][j] = newColor
	FloodFill(screen, i+1, j, color, newColor)
	FloodFill(screen, i-1, j, color, newColor)
	FloodFill(screen, i, j+1, color, newColor)
	FloodFill(screen, i, j-1, color, newColor)
}

func PaintFill(screen [][]int, r, c, newColor int) [][]int {
	if len(screen) == 0 || len(screen[0]) == 0 {
		panic("screen must be a non-empty 2D list")
	}
	if r < 0 || r >= len(screen) || c < 0 || c >= len(screen[0]) {
		panic("start coordinates out of bounds")
	}
	color := screen[r][c]
	if color != newColor {
		FloodFill(screen, r, c, color, newColor)
	}
	return screen
}

func RunQ810() {
	screen := [][]int{{1, 2, 5}, {2, 2, 4}, {2, 8, 6}}
	PaintFill(screen, 1, 1, 3)
	for _, row := range screen {
		fmt.Println(row)
	}
}
