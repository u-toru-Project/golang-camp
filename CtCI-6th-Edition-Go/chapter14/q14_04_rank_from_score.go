package chapter14

import "fmt"

type ScoreboardStore struct {
	scores map[string]int
}

func NewScoreboardStore() *ScoreboardStore {
	return &ScoreboardStore{scores: make(map[string]int)}
}

func (s *ScoreboardStore) UpsertScore(player string, points int) {
	s.scores[player] = points
}

func (s *ScoreboardStore) RankOf(player string) *int {
	points, ok := s.scores[player]
	if !ok {
		return nil
	}
	rank := 1
	for _, other := range s.scores {
		if other > points {
			rank++
		}
	}
	return &rank
}

func RunQ1404() {
	store := NewScoreboardStore()
	store.UpsertScore("a", 100)
	store.UpsertScore("b", 200)
	store.UpsertScore("c", 150)
	fmt.Printf("b=%d c=%d\n", *store.RankOf("b"), *store.RankOf("c"))
}
