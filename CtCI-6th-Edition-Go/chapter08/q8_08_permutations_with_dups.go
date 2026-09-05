package chapter08

import "fmt"

func PrintPerms(value string) []string {
	result := make([]string, 0)
	printPermsInner(BuildFreqTable(value), "", len(value), &result)
	return result
}

func BuildFreqTable(value string) map[rune]int {
	counts := map[rune]int{}
	for _, character := range value {
		counts[character]++
	}
	return counts
}

func printPermsInner(letterCountMap map[rune]int, prefix string, remaining int, result *[]string) {
	if remaining == 0 {
		*result = append(*result, prefix)
		return
	}
	keys := make([]rune, 0, len(letterCountMap))
	for character := range letterCountMap {
		keys = append(keys, character)
	}
	for _, character := range keys {
		count := letterCountMap[character]
		if count <= 0 {
			continue
		}
		letterCountMap[character] = count - 1
		printPermsInner(letterCountMap, prefix+string(character), remaining-1, result)
		letterCountMap[character] = count
	}
}

func RunQ808() {
	fmt.Println(PrintPerms("aaf"))
}
