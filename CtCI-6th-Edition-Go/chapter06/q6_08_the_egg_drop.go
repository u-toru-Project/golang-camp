package chapter06

import "fmt"

func MinDrops(floors, eggs int) int {
	if floors < 0 || eggs < 1 {
		panic("floors must be >= 0 and eggs >= 1")
	}
	if floors == 0 {
		return 0
	}
	if eggs == 1 {
		return floors
	}
	prev := make([]int, floors+1)
	for i := range prev {
		prev[i] = i
	}
	for egg := 2; egg <= eggs; egg++ {
		curr := make([]int, floors+1)
		for f := 1; f <= floors; f++ {
			best := floors + 1
			for drop := 1; drop <= f; drop++ {
				broken := curr[drop-1]
				intact := prev[f-drop]
				worst := 1 + broken
				if intact > broken {
					worst = 1 + intact
				}
				if worst < best {
					best = worst
				}
			}
			curr[f] = best
		}
		prev = curr
	}
	return prev[floors]
}

func MinDropsBinarySearch(floors, eggs int) int {
	return MinDrops(floors, eggs)
}

func RunQ608() {
	fmt.Printf("100 floors, 2 eggs: %d drops\n", MinDrops(100, 2))
}
