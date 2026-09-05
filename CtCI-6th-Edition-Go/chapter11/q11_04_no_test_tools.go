package chapter11

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func NormalizeName(first, last string) string {
	if first == "" || last == "" {
		panic("names required")
	}
	return toTitle(strings.TrimSpace(first)) + " " + toTitle(strings.TrimSpace(last))
}

func ShippingCost(weightKg, distanceKm float64) float64 {
	if weightKg <= 0 || distanceKm < 0 {
		panic("invalid input")
	}
	return roundAwayFromZero(weightKg*2.0+distanceKm*0.05, 2)
}

func toTitle(value string) string {
	lower := strings.ToLower(value)
	if lower == "" {
		return lower
	}
	runes := []rune(lower)
	runes[0] = unicode.ToTitle(runes[0])
	return string(runes)
}

func roundAwayFromZero(value float64, digits int) float64 {
	factor := math.Pow(10, float64(digits))
	if value >= 0 {
		return math.Floor(value*factor+0.5) / factor
	}
	return math.Ceil(value*factor-0.5) / factor
}

func RunQ1104() {
	fmt.Println(NormalizeName(" ada ", " lovelace "))
	fmt.Println(ShippingCost(10.0, 100.0))
}
