package chapter16

import "fmt"

func GetSizes(matrix [][]rune, typ rune) []int {
	if len(matrix) == 0 {
		return nil
	}
	h, w := len(matrix), len(matrix[0])
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	sizes := make([]int, 0)
	for row := range h {
		for col := range w {
			if matrix[row][col] == typ && !visited[row][col] {
				sizes = append(sizes, getPondSize(matrix, visited, row, col, typ))
			}
		}
	}
	return sizes
}

func getPondSize(matrix [][]rune, visited [][]bool, row, col int, typ rune) int {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || visited[row][col] || matrix[row][col] != typ {
		return 0
	}
	visited[row][col] = true
	size := 1
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			size += getPondSize(matrix, visited, row+dr, col+dc, typ)
		}
	}
	return size
}

func RunQ1619() {
	matrix := [][]rune{
		{'w', 'h', 'l', 'w'},
		{'w', 'l', 'w', '1'},
		{'l', 'l', 'w', 'l'},
		{'w', 'l', 'w', 'l'},
	}
	fmt.Printf("Pond sizes: %v\n", GetSizes(matrix, 'w'))
}
