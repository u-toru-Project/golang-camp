package chapter16

import (
	"fmt"
	"sort"
)

type Person struct {
	BirthYear int
	DeathYear int
}

type yearEvent struct {
	Year  int
	Delta int
}

func YearWithMostLiving(people []Person) int {
	if len(people) == 0 {
		panic("People must be non-empty.")
	}
	events := make([]yearEvent, 0, len(people)*2)
	for _, person := range people {
		if person.BirthYear > person.DeathYear {
			panic("Birth year cannot exceed death year.")
		}
		events = append(events, yearEvent{Year: person.BirthYear, Delta: 1})
		events = append(events, yearEvent{Year: person.DeathYear, Delta: -1})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].Year != events[j].Year {
			return events[i].Year < events[j].Year
		}
		return events[i].Delta > events[j].Delta
	})
	population := 0
	bestPopulation := 0
	bestYear := events[0].Year
	for _, ev := range events {
		population += ev.Delta
		if population > bestPopulation {
			bestPopulation = population
			bestYear = ev.Year
		}
	}
	return bestYear
}

func RunQ1610() {
	people := []Person{
		{12, 15}, {20, 90}, {10, 98}, {1, 72}, {10, 98},
		{23, 82}, {13, 98}, {90, 98}, {83, 99}, {75, 94},
	}
	fmt.Printf("Year with most living: %d\n", YearWithMostLiving(people))
}
