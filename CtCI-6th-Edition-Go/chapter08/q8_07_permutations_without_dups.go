package chapter08

import "fmt"

func GetPerms(value *string) []string {
	if value == nil {
		return nil
	}
	if len(*value) == 0 {
		return []string{""}
	}
	first := (*value)[0]
	rest := (*value)[1:]
	words := GetPerms(&rest)
	permutations := make([]string, 0)
	for _, word := range words {
		for index := 0; index <= len(word); index++ {
			permutations = append(permutations, InsertCharAt(word, rune(first), index))
		}
	}
	return permutations
}

func InsertCharAt(word string, character rune, index int) string {
	return word[:index] + string(character) + word[index:]
}

func GetPerms2(value string) []string {
	result := make([]string, 0)
	getPermsInner2("", value, &result)
	return result
}

func getPermsInner2(prefix, remainder string, result *[]string) {
	if len(remainder) == 0 {
		*result = append(*result, prefix)
		return
	}
	for i := 0; i < len(remainder); i++ {
		before := remainder[:i]
		after := remainder[i+1:]
		getPermsInner2(prefix+string(remainder[i]), before+after, result)
	}
}

func RunQ807() {
	s := "str"
	fmt.Println(GetPerms(&s))
	fmt.Println(GetPerms2("str"))
}
