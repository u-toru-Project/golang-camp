package chapter14

import "fmt"

type InventoryStore struct {
	qty map[string]int
}

func NewInventoryStore() *InventoryStore {
	return &InventoryStore{qty: make(map[string]int)}
}

func (s *InventoryStore) AdjustInventory(sku string, delta int) int {
	current := s.qty[sku]
	newQty := current + delta
	if newQty < 0 {
		panic("insufficient stock")
	}
	s.qty[sku] = newQty
	return newQty
}

func RunQ1406() {
	store := NewInventoryStore()
	fmt.Println(store.AdjustInventory("A1", 10))
	fmt.Println(store.AdjustInventory("A1", -3))
}
