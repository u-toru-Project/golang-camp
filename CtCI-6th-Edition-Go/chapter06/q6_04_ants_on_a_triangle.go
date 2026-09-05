package chapter06

import (
	"fmt"
	"math"
)

func CollisionProbability(numAnts int) float64 {
	if numAnts < 2 {
		panic("num_ants must be at least 2")
	}
	allCases := math.Pow(2, float64(numAnts))
	noCollisionCases := 2.0
	collisionCases := allCases - noCollisionCases
	if allCases == 0 {
		return 0.0
	}
	return collisionCases / allCases
}

func CollisionProbabilityClosedForm(numAnts int) float64 {
	denominator := math.Pow(2, float64(numAnts))
	if denominator == 0 {
		return 0.0
	}
	return (denominator - 2) / denominator
}

func RunQ604() {
	fmt.Printf("Triangle collision probability: %v\n", CollisionProbability(3))
}
