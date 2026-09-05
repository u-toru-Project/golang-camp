package chapter17

import "fmt"

func CountBabyNames(nameCounts map[string]int, synonyms [][2]string) map[string]int {
	parent := make(map[string]string, len(nameCounts))
	for name := range nameCounts {
		parent[name] = name
	}
	for _, pair := range synonyms {
		unionNames(parent, pair[0], pair[1])
	}
	totals := make(map[string]int)
	for name := range parent {
		root := findName(parent, name)
		totals[root] += nameCounts[name]
	}
	return totals
}

func findName(parent map[string]string, name string) string {
	if _, ok := parent[name]; !ok {
		panic("Unknown name: " + name)
	}
	if parent[name] != name {
		parent[name] = findName(parent, parent[name])
	}
	return parent[name]
}

func unionNames(parent map[string]string, first, second string) {
	rootFirst := findName(parent, first)
	rootSecond := findName(parent, second)
	if rootFirst != rootSecond {
		parent[rootFirst] = rootSecond
	}
}

func RunQ1707() {
	nameCounts := map[string]int{
		"john": 10, "jon": 3, "davis": 2, "kari": 3, "johny": 11,
		"carlton": 8, "carleton": 2, "jonathan": 9, "carrie": 5,
	}
	synonyms := [][2]string{{"jonathan", "john"}, {"jon", "johny"}, {"johny", "john"}, {"kari", "carrie"}, {"carleton", "carlton"}}
	fmt.Println(CountBabyNames(nameCounts, synonyms))
}
