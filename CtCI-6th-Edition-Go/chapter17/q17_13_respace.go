package chapter17

import "strings"

import "fmt"

func Respace(text string, dictionary []string) string {
	wordSet := make(map[string]struct{}, len(dictionary))
	for _, w := range dictionary {
		wordSet[w] = struct{}{}
	}
	memo := make(map[int]respaceResult)
	result := respaceHelper(text, 0, wordSet, memo)
	return applySplits(text, result.splits)
}

type respaceResult struct {
	splits  []int
	invalid int
}

func respaceHelper(text string, start int, wordSet map[string]struct{}, memo map[int]respaceResult) respaceResult {
	if cached, ok := memo[start]; ok {
		return cached
	}
	if start == len(text) {
		return respaceResult{}
	}
	bestInvalid := len(text)
	bestSplits := []int{}
	var piece strings.Builder
	for end := start; end < len(text); end++ {
		piece.WriteString(string(text[end]))
		invalid := len(piece.String())
		if _, ok := wordSet[piece.String()]; ok {
			invalid = 0
		}
		if invalid >= bestInvalid {
			continue
		}
		rest := respaceHelper(text, end+1, wordSet, memo)
		totalInvalid := invalid + rest.invalid
		if totalInvalid < bestInvalid {
			bestInvalid = totalInvalid
			splits := append([]int(nil), rest.splits...)
			if invalid == 0 && end < len(text)-1 {
				splits = append(splits, end)
			}
			bestSplits = splits
		}
	}
	memo[start] = respaceResult{bestSplits, bestInvalid}
	return memo[start]
}

func applySplits(text string, splitAfter []int) string {
	splitSet := make(map[int]struct{}, len(splitAfter))
	for _, s := range splitAfter {
		splitSet[s] = struct{}{}
	}
	out := make([]byte, 0, len(text)+len(splitSet))
	for i := 0; i < len(text); i++ {
		out = append(out, text[i])
		if _, ok := splitSet[i]; ok {
			out = append(out, ' ')
		}
	}
	return string(out)
}

func RunQ1713() {
	words := []string{"help", "tech", "tips", "boot", "my", "reboot", "need", "to", "linus"}
	fmt.Println(Respace("helplinustechtipsineedtorebootmypc", words))
	fmt.Println(Respace("abc", []string{"a", "bc"}))
}
