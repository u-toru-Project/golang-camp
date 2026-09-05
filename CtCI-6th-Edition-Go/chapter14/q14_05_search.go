package chapter14

import (
	"fmt"
	"strings"
)

type DocumentHit struct {
	Id    int
	Title string
}

type document struct {
	id    int
	title string
	body  string
}

type DocumentSearchStore struct {
	nextID    int
	documents []document
}

func NewDocumentSearchStore() *DocumentSearchStore {
	return &DocumentSearchStore{nextID: 1}
}

func (s *DocumentSearchStore) AddDocument(title, body string) {
	s.documents = append(s.documents, document{id: s.nextID, title: title, body: body})
	s.nextID++
}

func (s *DocumentSearchStore) SearchDocuments(query string) []DocumentHit {
	if query == "" {
		panic("query required")
	}
	needle := strings.ToLower(query)
	hits := make([]DocumentHit, 0)
	for _, doc := range s.documents {
		if strings.Contains(strings.ToLower(doc.title), needle) || strings.Contains(strings.ToLower(doc.body), needle) {
			hits = append(hits, DocumentHit{Id: doc.id, Title: doc.title})
		}
	}
	return hits
}

func RunQ1405() {
	store := NewDocumentSearchStore()
	store.AddDocument("Python Tips", "use pytest")
	store.AddDocument("SQL Guide", "joins and indexes")
	hits := store.SearchDocuments("pytest")
	parts := make([]string, 0, len(hits))
	for _, hit := range hits {
		parts = append(parts, fmt.Sprintf("%d:%s", hit.Id, hit.Title))
	}
	fmt.Println(strings.Join(parts, ", "))
}
