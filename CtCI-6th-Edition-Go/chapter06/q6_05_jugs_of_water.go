package chapter06

import "fmt"

func MeasureableAmounts(capA, capB int) map[int]struct{} {
	if capA < 1 || capB < 1 {
		panic("capacities must be positive")
	}
	g := gcd(capA, capB)
	amounts := make(map[int]struct{})
	maxCap := max(capB, capA)
	for i := g; i <= maxCap; i += g {
		amounts[i] = struct{}{}
	}
	return amounts
}

func CanMeasure(capA, capB, target int) bool {
	if target < 0 {
		panic("target must be non-negative")
	}
	maxCap := max(capB, capA)
	if target > maxCap {
		return false
	}
	if target == 0 {
		return true
	}
	_, ok := MeasureableAmounts(capA, capB)[target]
	return ok
}

func MinPoursBfs(capA, capB, target int) *int {
	if !CanMeasure(capA, capB, target) {
		return nil
	}
	type state struct{ a, b, steps int }
	queue := []state{{0, 0, 0}}
	seen := map[[2]int]struct{}{{0, 0}: {}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.a == target || cur.b == target {
			steps := cur.steps
			return &steps
		}
		for _, next := range jugNeighbors(cur.a, cur.b, capA, capB) {
			key := [2]int{next[0], next[1]}
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			queue = append(queue, state{next[0], next[1], cur.steps + 1})
		}
	}
	return nil
}

func MeasureFourLiters() *int {
	return MinPoursBfs(3, 5, 4)
}

func jugNeighbors(a, b, capA, capB int) [][2]int {
	pourAToB := min(capB-b, a)
	pourBToA := min(capA-a, b)
	return [][2]int{
		{capA, b},
		{a, capB},
		{0, b},
		{a, 0},
		{a - pourAToB, b + pourAToB},
		{a + pourBToA, b - pourBToA},
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func RunQ605() {
	fmt.Printf("Can measure 4L with 3 and 5: %t\n", CanMeasure(3, 5, 4))
	fmt.Printf("Min pours: %v\n", *MeasureFourLiters())
}
