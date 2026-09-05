package chapter06

import "fmt"

func DaysUntilBlueEyedLeave(numBlueEyed int, guruAnnounced bool) int {
	if numBlueEyed < 0 {
		panic("num_blue_eyed must be non-negative")
	}
	if numBlueEyed == 0 || !guruAnnounced {
		return 0
	}
	return numBlueEyed
}

func SimulateLeavingDay(eyeColors []string) int {
	if eyeColors == nil {
		panic("eyeColors is nil")
	}
	if len(eyeColors) == 0 {
		panic("population must be non-empty")
	}
	blue := 0
	for _, color := range eyeColors {
		if color != "B" && color != "R" {
			panic("colors must be 'B' or 'R'")
		}
		if color == "B" {
			blue++
		}
	}
	return DaysUntilBlueEyedLeave(blue, true)
}

func RunQ606() {
	fmt.Printf("5 blue-eyed leave on day %d\n", DaysUntilBlueEyedLeave(5, true))
	fmt.Printf("No guru: %d\n", DaysUntilBlueEyedLeave(10, false))
}
