package chapter16

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) equal(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

type line03 struct {
	Slope      *float64
	Intercept  float64
	IsVertical bool
}

func line03Equal(a, b line03) bool {
	if a.IsVertical != b.IsVertical {
		return false
	}
	if a.IsVertical {
		return a.Intercept == b.Intercept
	}
	if a.Slope == nil || b.Slope == nil {
		return false
	}
	return *a.Slope == *b.Slope && a.Intercept == b.Intercept
}

func SegmentIntersection(a1, a2, b1, b2 Point) *Point {
	line1 := lineThrough03(a1, a2)
	line2 := lineThrough03(b1, b2)
	if line03Equal(line1, line2) || (line1.IsVertical && line2.IsVertical && line1.Intercept == line2.Intercept) {
		if hit := overlapEndpoint(a1, a2, b1, b2); hit != nil {
			return hit
		}
		return overlapEndpoint(b1, b2, a1, a2)
	}
	hit := intersectLines03(line1, line2)
	if hit == nil {
		return nil
	}
	if onSegment03(*hit, a1, a2) && onSegment03(*hit, b1, b2) {
		return hit
	}
	return nil
}

func lineThrough03(p1, p2 Point) line03 {
	if p1.equal(p2) {
		panic("Segment endpoints must differ.")
	}
	if p1.X == p2.X {
		return line03{Intercept: p1.X, IsVertical: true}
	}
	slope := (p1.Y - p2.Y) / (p1.X - p2.X)
	return line03{Slope: &slope, Intercept: p1.Y - slope*p1.X}
}

func intersectLines03(line1, line2 line03) *Point {
	if line1.IsVertical && line2.IsVertical {
		if line1.Intercept != line2.Intercept {
			return nil
		}
		return &Point{X: line1.Intercept, Y: 0}
	}
	if line1.IsVertical {
		x := line1.Intercept
		return &Point{X: x, Y: *line2.Slope*x + line2.Intercept}
	}
	if line2.IsVertical {
		x := line2.Intercept
		return &Point{X: x, Y: *line1.Slope*x + line1.Intercept}
	}
	if *line1.Slope == *line2.Slope {
		if line1.Intercept != line2.Intercept {
			return nil
		}
		return &Point{X: 0, Y: line1.Intercept}
	}
	x := (line2.Intercept - line1.Intercept) / (*line1.Slope - *line2.Slope)
	return &Point{X: x, Y: *line1.Slope*x + line1.Intercept}
}

func onSegment03(point, start, end Point) bool {
	const epsilon = 1e-9
	minX := math.Min(start.X, end.X) - epsilon
	maxX := math.Max(start.X, end.X) + epsilon
	minY := math.Min(start.Y, end.Y) - epsilon
	maxY := math.Max(start.Y, end.Y) + epsilon
	return minX <= point.X && point.X <= maxX && minY <= point.Y && point.Y <= maxY
}

func overlapEndpoint(start, end, otherStart, otherEnd Point) *Point {
	if onSegment03(start, otherStart, otherEnd) {
		return &Point{X: start.X, Y: start.Y}
	}
	if onSegment03(end, otherStart, otherEnd) {
		return &Point{X: end.X, Y: end.Y}
	}
	return nil
}

func RunQ1603() {
	hit := SegmentIntersection(Point{10, 10}, Point{20, 20}, Point{20, 10}, Point{10, 20})
	if hit == nil {
		fmt.Println("No intersection")
		return
	}
	fmt.Printf("(%v, %v)\n", hit.X, hit.Y)
}
