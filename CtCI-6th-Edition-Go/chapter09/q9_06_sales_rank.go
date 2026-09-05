package chapter09

import (
	"fmt"
	"sort"
)

type SaleCount struct {
	ProductId string
	Count     int
}

type SalesRank struct {
	k      int
	counts map[string]int
}

func NewSalesRank(k int) *SalesRank {
	if k < 1 {
		panic("k must be positive")
	}
	return &SalesRank{k: k, counts: map[string]int{}}
}

func (s *SalesRank) RecordSale(productId string, quantity ...int) {
	q := 1
	if len(quantity) > 0 {
		q = quantity[0]
	}
	if q < 1 {
		panic("quantity must be positive")
	}
	s.counts[productId] += q
}

func (s *SalesRank) TopK() []SaleCount {
	items := make([]SaleCount, 0, len(s.counts))
	for id, count := range s.counts {
		items = append(items, SaleCount{id, count})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Count != items[j].Count {
			return items[i].Count > items[j].Count
		}
		return items[i].ProductId > items[j].ProductId
	})
	if s.k < len(items) {
		items = items[:s.k]
	}
	return items
}

func RunQ906() {
	rank := NewSalesRank(2)
	rank.RecordSale("p1", 5)
	rank.RecordSale("p2", 10)
	rank.RecordSale("p3", 7)
	fmt.Println(rank.TopK())
}
