package chapter17

import (
	"fmt"
	"math"
	"sort"
)

type Document struct {
	Id     int
	Tokens []int
}

type SimilarityPair struct {
	Id1        int
	Id2        int
	Similarity float64
}

func SparseSimilarity(documents []Document) []SimilarityPair {
	docMap := make(map[int][]int)
	for _, doc := range documents {
		docMap[doc.Id] = doc.Tokens
	}
	inverted := make(map[int]map[int]struct{})
	for docID, tokens := range docMap {
		for _, token := range tokens {
			owners := inverted[token]
			if owners == nil {
				owners = make(map[int]struct{})
				inverted[token] = owners
			}
			owners[docID] = struct{}{}
		}
	}
	overlap := make(map[[2]int]int)
	for docID, tokens := range docMap {
		for _, token := range tokens {
			for otherID := range inverted[token] {
				if otherID <= docID {
					continue
				}
				pair := [2]int{docID, otherID}
				overlap[pair]++
			}
		}
	}
	keys := make([][2]int, 0, len(overlap))
	for k := range overlap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i][0] != keys[j][0] {
			return keys[i][0] < keys[j][0]
		}
		return keys[i][1] < keys[j][1]
	})
	results := make([]SimilarityPair, 0, len(keys))
	for _, pair := range keys {
		common := overlap[pair]
		unionSize := len(docMap[pair[0]]) + len(docMap[pair[1]]) - common
		similarity := 0.0
		if unionSize != 0 {
			similarity = float64(common) / float64(unionSize)
		}
		results = append(results, SimilarityPair{Id1: pair[0], Id2: pair[1], Similarity: similarity})
	}
	return results
}

func SparseDotSimilarity(vecA, vecB map[int]int) float64 {
	if len(vecA) == 0 || len(vecB) == 0 {
		return 0.0
	}
	var dot int64
	for key, weight := range vecA {
		if other, ok := vecB[key]; ok {
			dot += int64(weight) * int64(other)
		}
	}
	normA := 0.0
	for _, value := range vecA {
		normA += float64(int64(value) * int64(value))
	}
	normB := 0.0
	for _, value := range vecB {
		normB += float64(int64(value) * int64(value))
	}
	normA = math.Sqrt(normA)
	normB = math.Sqrt(normB)
	if normA == 0 || normB == 0 {
		return 0.0
	}
	return float64(dot) / (normA * normB)
}

func RunQ1726() {
	docs := []Document{
		{13, []int{14, 15, 100, 9, 3}},
		{16, []int{32, 1, 9, 3, 5}},
		{19, []int{15, 29, 2, 6, 8, 7}},
		{24, []int{7, 10}},
	}
	for _, pair := range SparseSimilarity(docs) {
		fmt.Printf("%d, %d: %v\n", pair.Id1, pair.Id2, pair.Similarity)
	}
}
