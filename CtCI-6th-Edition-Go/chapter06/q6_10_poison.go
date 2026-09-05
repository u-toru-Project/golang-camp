package chapter06

import "fmt"

const DaysForResult = 7

type testStrip struct {
	hasPoison   bool
	dayPoisoned *int
}

type World struct {
	NumBottles        int
	NumTestStrips     int
	day               int
	poisonedBottleNum int
	testStrips        []testStrip
}

func NewWorld(numTestStrips, numBottles, poisonedBottleNum int) *World {
	if numTestStrips < 1 || numBottles < 1 {
		panic("counts must be positive")
	}
	if poisonedBottleNum < 0 || poisonedBottleNum >= numBottles {
		panic("poisoned_bottle_num out of range")
	}
	if numBottles > (1 << numTestStrips) {
		panic("not enough strips to identify all bottles")
	}
	strips := make([]testStrip, numTestStrips)
	return &World{
		NumBottles:        numBottles,
		NumTestStrips:     numTestStrips,
		poisonedBottleNum: poisonedBottleNum,
		testStrips:        strips,
	}
}

func (w *World) Day() int { return w.day }

func (w *World) SetDay(value int) {
	if value < w.day {
		panic("day cannot be decreased")
	}
	w.day = value
}

func (w *World) AddDrop(bottleNum, testStripNum int) {
	if bottleNum < 0 || bottleNum >= w.NumBottles {
		panic("bottle_num out of range")
	}
	if testStripNum < 0 || testStripNum >= w.NumTestStrips {
		panic("test_strip_num out of range")
	}
	strip := &w.testStrips[testStripNum]
	if bottleNum == w.poisonedBottleNum && !strip.hasPoison {
		strip.hasPoison = true
		day := w.day
		strip.dayPoisoned = &day
	}
}

func (w *World) PositiveTestStrips() []int {
	result := make([]int, 0)
	for i, strip := range w.testStrips {
		if strip.hasPoison && strip.dayPoisoned != nil && w.day-*strip.dayPoisoned >= DaysForResult {
			result = append(result, i)
		}
	}
	return result
}

func FindPoison(world *World) int {
	if world == nil {
		panic("world is nil")
	}
	for i := 0; i < world.NumBottles; i++ {
		for j := 0; j < world.NumTestStrips; j++ {
			if i&(1<<j) != 0 {
				world.AddDrop(i, j)
			}
		}
	}
	world.SetDay(world.Day() + DaysForResult)
	bottle := 0
	for _, strip := range world.PositiveTestStrips() {
		bottle |= 1 << strip
	}
	return bottle
}

func RunQ610() {
	world := NewWorld(10, 1000, 42)
	fmt.Printf("Poisoned bottle: %d\n", FindPoison(world))
}
