package chapter08

import "fmt"

type GridPoint struct {
	Row int
	Col int
}

func GetPath(maze [][]bool) []GridPoint {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return nil
	}
	path := make([]GridPoint, 0)
	if isPath(maze, len(maze)-1, len(maze[0])-1, &path) {
		return path
	}
	return nil
}

func isPath(maze [][]bool, row, col int, path *[]GridPoint) bool {
	if col < 0 || row < 0 || !maze[row][col] {
		return false
	}
	isAtOrigin := row == 0 && col == 0
	if isAtOrigin || isPath(maze, row, col-1, path) || isPath(maze, row-1, col, path) {
		*path = append(*path, GridPoint{row, col})
		return true
	}
	return false
}

func GetPathMemoized(maze [][]bool) []GridPoint {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return nil
	}
	path := make([]GridPoint, 0)
	failed := map[GridPoint]struct{}{}
	if isPathMemoized(maze, len(maze)-1, len(maze[0])-1, &path, failed) {
		return path
	}
	return nil
}

func isPathMemoized(maze [][]bool, row, col int, path *[]GridPoint, failed map[GridPoint]struct{}) bool {
	if col < 0 || row < 0 || !maze[row][col] {
		return false
	}
	point := GridPoint{row, col}
	if _, ok := failed[point]; ok {
		return false
	}
	isAtOrigin := row == 0 && col == 0
	if isAtOrigin || isPathMemoized(maze, row, col-1, path, failed) || isPathMemoized(maze, row-1, col, path, failed) {
		*path = append(*path, point)
		return true
	}
	failed[point] = struct{}{}
	return false
}

func RunQ802() {
	maze := [][]bool{{true, true}, {false, true}}
	fmt.Printf("path=%v memo=%v\n", GetPath(maze), GetPathMemoized(maze))
}
