package chapter14

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type LanguageStore struct {
	users     map[int]string
	languages map[int][]string
}

func NewLanguageStore() *LanguageStore {
	return &LanguageStore{
		users:     make(map[int]string),
		languages: make(map[int][]string),
	}
}

func (s *LanguageStore) AddUser(id int, name string) {
	s.users[id] = name
}

func (s *LanguageStore) SetLanguages(userID int, langs []string) {
	copied := make([]string, len(langs))
	copy(copied, langs)
	s.languages[userID] = copied
}

func (s *LanguageStore) UsersSpeaking(langCode string) []string {
	names := make([]string, 0)
	for userID, langs := range s.languages {
		if slices.Contains(langs, langCode) {
			names = append(names, s.users[userID])
		}
	}
	sort.Strings(names)
	return names
}

func RunQ1402() {
	store := NewLanguageStore()
	store.AddUser(1, "Alice")
	store.AddUser(2, "Bob")
	store.SetLanguages(1, []string{"en", "fr"})
	store.SetLanguages(2, []string{"en"})
	fmt.Printf("en: %s\n", strings.Join(store.UsersSpeaking("en"), ", "))
	fmt.Printf("fr: %s\n", strings.Join(store.UsersSpeaking("fr"), ", "))
}
