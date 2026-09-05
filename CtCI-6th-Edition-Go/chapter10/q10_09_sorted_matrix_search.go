package chapter10

import "fmt"

type Coordinate struct {
	Row    int
	Column int
}

func NewCoordinate(row, column int) *Coordinate {
	return &Coordinate{Row: row, Column: column}
}

func (c *Coordinate) InBounds(matrix [][]int) bool {
	return len(matrix) > 0 && len(matrix[0]) > 0 && c.Row >= 0 && c.Column >= 0 && c.Row < len(matrix) && c.Column < len(matrix[0])
}

func (c *Coordinate) IsBefore(other *Coordinate) bool {
	return c.Row <= other.Row && c.Column <= other.Column
}

func (c *Coordinate) Clone() *Coordinate {
	return NewCoordinate(c.Row, c.Column)
}

func (c *Coordinate) SetToAverage(min, max *Coordinate) {
	c.Row = (min.Row + max.Row) / 2
	c.Column = (min.Column + max.Column) / 2
}

func FindElement(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := 0
	column := len(matrix[0]) - 1
	for row < len(matrix) && column >= 0 {
		if matrix[row][column] == target {
			return true
		}
		if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}

func FindElement2(matrix [][]int, target int) *Coordinate {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	origin := NewCoordinate(0, 0)
	destination := NewCoordinate(len(matrix)-1, len(matrix[0])-1)
	return findElement2(matrix, origin, destination, target)
}

func findElement2(matrix [][]int, origin, destination *Coordinate, target int) *Coordinate {
	if !origin.InBounds(matrix) || !destination.InBounds(matrix) {
		return nil
	}
	if matrix[origin.Row][origin.Column] == target {
		return origin
	}
	if !origin.IsBefore(destination) {
		return nil
	}
	start := origin.Clone()
	diagonalDistance := min(destination.Column-origin.Column, destination.Row-origin.Row)
	end := NewCoordinate(start.Row+diagonalDistance, start.Column+diagonalDistance)
	pivot := NewCoordinate(0, 0)
	for start.IsBefore(end) {
		pivot.SetToAverage(start, end)
		value := matrix[pivot.Row][pivot.Column]
		if value == target {
			return NewCoordinate(pivot.Row, pivot.Column)
		}
		if target > value {
			start.Row = pivot.Row + 1
			start.Column = pivot.Column + 1
		} else {
			end.Row = pivot.Row - 1
			end.Column = pivot.Column - 1
		}
	}
	if start.InBounds(matrix) && matrix[start.Row][start.Column] == target {
		return start
	}
	return partitionAndSearch(matrix, origin, destination, start, target)
}

func partitionAndSearch(matrix [][]int, origin, destination, pivot *Coordinate, target int) *Coordinate {
	lowerLeftOrigin := NewCoordinate(pivot.Row, origin.Column)
	lowerLeftDestination := NewCoordinate(destination.Row, pivot.Column-1)
	upperRightOrigin := NewCoordinate(origin.Row, pivot.Column)
	upperRightDestination := NewCoordinate(pivot.Row-1, destination.Column)
	if found := findElement2(matrix, lowerLeftOrigin, lowerLeftDestination, target); found != nil {
		return found
	}
	return findElement2(matrix, upperRightOrigin, upperRightDestination, target)
}

func RunQ1009() {
	matrix := [][]int{
		{15, 30, 50, 70, 73},
		{35, 40, 100, 102, 120},
		{36, 42, 105, 110, 125},
		{46, 51, 106, 111, 130},
		{48, 55, 109, 140, 150},
	}
	for _, target := range []int{15, 40, 110, 150, 16, 151} {
		coord := FindElement2(matrix, target)
		location := "not found"
		if coord != nil {
			location = fmt.Sprintf("(%d, %d)", coord.Row, coord.Column)
		}
		fmt.Printf("%d: %t / %s\n", target, FindElement(matrix, target), location)
	}
}
