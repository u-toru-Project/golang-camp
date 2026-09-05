package chapter16

import (
	"fmt"
	"maps"
	"math"
)

type Line struct {
	Slope      *float64
	Intercept  float64
	IsVertical bool
}

func lineKey(a Line) string {
	if a.IsVertical {
		return fmt.Sprintf("v:%v", a.Intercept)
	}
	return fmt.Sprintf("s:%v:%v", *a.Slope, a.Intercept)
}

func BestLine(points []Point) (Line, int) {
	if len(points) < 2 {
		panic("Need at least two points.")
	}
	counts := make(map[string]int)
	lines := make(map[string]Line)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			if points[i].equal(points[j]) {
				continue
			}
			key := lineKey14(points[i], points[j])
			k := lineKey(key)
			counts[k]++
			lines[k] = key
		}
	}
	merged := mergeLines(counts, lines, 1e-4)
	bestLine := Line{}
	bestCount := -1
	for k, v := range merged {
		if v > bestCount {
			bestCount = v
			bestLine = lines[k]
		}
	}
	return bestLine, bestCount + 1
}

func lineKey14(p1, p2 Point) Line {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	if dx == 0 {
		return Line{Intercept: p1.X, IsVertical: true}
	}
	slope := dy / dx
	return Line{Slope: &slope, Intercept: p1.Y - slope*p1.X}
}

func mergeLines(counts map[string]int, lines map[string]Line, epsilon float64) map[string]int {
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	merged := make(map[string]int)
	maps.Copy(merged, counts)
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			first := lines[keys[i]]
			second := lines[keys[j]]
			if first.IsVertical != second.IsVertical {
				continue
			}
			if first.IsVertical {
				if math.Abs(first.Intercept-second.Intercept) <= epsilon {
					merged[keys[i]] = merged[keys[i]] + merged[keys[j]]
					delete(merged, keys[j])
				}
			} else if first.Slope != nil && second.Slope != nil &&
				math.Abs(*first.Slope-*second.Slope) <= epsilon &&
				math.Abs(first.Intercept-second.Intercept) <= epsilon {
				merged[keys[i]] = merged[keys[i]] + merged[keys[j]]
				delete(merged, keys[j])
			}
		}
	}
	return merged
}

func RunQ1614() {
	points := []Point{{1, 1}, {4, 1}, {7, 1}, {8, 3}, {3, 3}}
	line, count := BestLine(points)
	fmt.Printf("vertical=%t slope=%v intercept=%v count=%d\n", line.IsVertical, line.Slope, line.Intercept, count)
}
