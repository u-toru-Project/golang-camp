package chapter07

import (
	"fmt"
	"hash/fnv"
)

type pair[K comparable, V any] struct {
	key   K
	value V
}

type HashMap[K comparable, V any] struct {
	buckets  [][]pair[K, V]
	capacity int
	size     int
}

func NewHashMap[K comparable, V any](capacity ...int) *HashMap[K, V] {
	cap := 16
	if len(capacity) > 0 {
		cap = capacity[0]
	}
	if cap < 1 {
		panic("capacity must be positive")
	}
	return &HashMap[K, V]{buckets: make([][]pair[K, V], cap), capacity: cap}
}

func (m *HashMap[K, V]) Count() int { return m.size }

func (m *HashMap[K, V]) Put(key K, value V) {
	idx := m.index(key)
	bucket := m.buckets[idx]
	for i, p := range bucket {
		if p.key == key {
			bucket[i].value = value
			m.buckets[idx] = bucket
			return
		}
	}
	m.buckets[idx] = append(bucket, pair[K, V]{key, value})
	m.size++
	if m.size > m.capacity*2 {
		m.rehash()
	}
}

func (m *HashMap[K, V]) Get(key K) V {
	var zero V
	for _, p := range m.buckets[m.index(key)] {
		if p.key == key {
			return p.value
		}
	}
	return zero
}

func (m *HashMap[K, V]) Delete(key K) bool {
	idx := m.index(key)
	bucket := m.buckets[idx]
	for i, p := range bucket {
		if p.key == key {
			m.buckets[idx] = append(bucket[:i], bucket[i+1:]...)
			m.size--
			return true
		}
	}
	return false
}

func (m *HashMap[K, V]) index(key K) int {
	h := fnv.New32a()
	fmt.Fprintf(h, "%v", key)
	return int(h.Sum32()&0x7fffffff) % m.capacity
}

func (m *HashMap[K, V]) rehash() {
	old := m.buckets
	m.capacity *= 2
	m.size = 0
	m.buckets = make([][]pair[K, V], m.capacity)
	for _, bucket := range old {
		for _, p := range bucket {
			m.Put(p.key, p.value)
		}
	}
}

func RunQ712() {
	m := NewHashMap[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	fmt.Println(m.Get("a"))
}
