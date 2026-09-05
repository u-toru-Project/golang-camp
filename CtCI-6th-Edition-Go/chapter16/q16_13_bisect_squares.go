package chapter16

import (
	"fmt"
	"math"
	"sort"
)

type Square struct {
	TopLeft, TopRight, BottomLeft, BottomRight Point
}

func (s Square) Center() Point {
	return Point{X: (s.TopLeft.X + s.BottomRight.X) / 2, Y: (s.TopLeft.Y + s.BottomRight.Y) / 2}
}

func (s Square) Sides() [][2]Point {
	return [][2]Point{
		{s.TopLeft, s.TopRight},
		{s.TopLeft, s.BottomLeft},
		{s.BottomRight, s.BottomLeft},
		{s.BottomRight, s.TopRight},
	}
}

type Segment struct {
	Start, End Point
}

func SquareCutSegment(first, second Square) Segment {
	center1 := first.Center()
	center2 := second.Center()
	if center1.equal(center2) {
		panic("Squares share the same center.")
	}
	midLine := lineThrough13(center1, center2)
	hits := make([]Point, 0)
	for _, square := range []Square{first, second} {
		for _, side := range square.Sides() {
			point := intersect13(lineThrough13(side[0], side[1]), midLine)
			if point != nil && onSegment13(*point, side[0], side[1]) {
				hits = append(hits, *point)
			}
		}
	}
	if len(hits) < 2 {
		panic("Could not find cut segment.")
	}
	sort.Slice(hits, func(i, j int) bool {
		if hits[i].X != hits[j].X {
			return hits[i].X < hits[j].X
		}
		return hits[i].Y < hits[j].Y
	})
	return Segment{Start: hits[0], End: hits[len(hits)-1]}
}

type line13 struct {
	Slope      *float64
	Intercept  float64
	IsVertical bool
}

func lineThrough13(p1, p2 Point) line13 {
	if p1.equal(p2) {
		panic("Points must differ.")
	}
	if p1.X == p2.X {
		return line13{Intercept: p1.X, IsVertical: true}
	}
	slope := (p1.Y - p2.Y) / (p1.X - p2.X)
	return line13{Slope: &slope, Intercept: p1.Y - slope*p1.X}
}

func intersect13(line1, line2 line13) *Point {
	if line1.IsVertical && line2.IsVertical {
		if line1.Intercept != line2.Intercept {
			return nil
		}
		return &Point{X: line1.Intercept, Y: 0}
	}
	if line1.IsVertical {
		return &Point{X: line1.Intercept, Y: *line2.Slope*line1.Intercept + line2.Intercept}
	}
	if line2.IsVertical {
		return &Point{X: line2.Intercept, Y: *line1.Slope*line2.Intercept + line1.Intercept}
	}
	if *line1.Slope == *line2.Slope {
		return nil
	}
	x := (line2.Intercept - line1.Intercept) / (*line1.Slope - *line2.Slope)
	return &Point{X: x, Y: *line1.Slope*x + line1.Intercept}
}

func onSegment13(point, start, end Point) bool {
	const epsilon = 1e-6
	minX := math.Min(start.X, end.X) - epsilon
	maxX := math.Max(start.X, end.X) + epsilon
	minY := math.Min(start.Y, end.Y) - epsilon
	maxY := math.Max(start.Y, end.Y) + epsilon
	return minX <= point.X && point.X <= maxX && minY <= point.Y && point.Y <= maxY
}

func RunQ1613() {
	first := Square{Point{2, 5}, Point{6, 5}, Point{2, 1}, Point{6, 1}}
	second := Square{Point{7, 8}, Point{9, 8}, Point{7, 6}, Point{9, 6}}
	seg := SquareCutSegment(first, second)
	fmt.Printf("(%v, %v) -> (%v, %v)\n", seg.Start.X, seg.Start.Y, seg.End.X, seg.End.Y)
}
