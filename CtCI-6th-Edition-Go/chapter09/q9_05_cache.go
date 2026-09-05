package chapter09

import "fmt"

type lruNode[K comparable, V any] struct {
	key   K
	value V
	prev  *lruNode[K, V]
	next  *lruNode[K, V]
}

type LruCache[K comparable, V any] struct {
	capacity int
	items    map[K]*lruNode[K, V]
	head     *lruNode[K, V]
	tail     *lruNode[K, V]
}

func NewLruCache[K comparable, V any](capacity int) *LruCache[K, V] {
	if capacity < 1 {
		panic("capacity must be positive")
	}
	return &LruCache[K, V]{capacity: capacity, items: map[K]*lruNode[K, V]{}}
}

func (c *LruCache[K, V]) Count() int { return len(c.items) }

func (c *LruCache[K, V]) TryGet(key K) (V, bool) {
	var zero V
	node, ok := c.items[key]
	if !ok {
		return zero, false
	}
	c.moveToEnd(node)
	return node.value, true
}

func (c *LruCache[K, V]) Put(key K, value V) {
	if existing, ok := c.items[key]; ok {
		existing.value = value
		c.moveToEnd(existing)
		return
	}
	node := &lruNode[K, V]{key: key, value: value}
	c.items[key] = node
	c.append(node)
	if len(c.items) > c.capacity {
		oldest := c.head
		c.remove(oldest)
		delete(c.items, oldest.key)
	}
}

func (c *LruCache[K, V]) append(node *lruNode[K, V]) {
	if c.tail == nil {
		c.head = node
		c.tail = node
		return
	}
	node.prev = c.tail
	c.tail.next = node
	c.tail = node
}

func (c *LruCache[K, V]) remove(node *lruNode[K, V]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		c.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		c.tail = node.prev
	}
	node.prev = nil
	node.next = nil
}

func (c *LruCache[K, V]) moveToEnd(node *lruNode[K, V]) {
	c.remove(node)
	c.append(node)
}

func RunQ905() {
	cache := NewLruCache[string, int](2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.TryGet("a")
	cache.Put("c", 3)
	_, hasB := cache.TryGet("b")
	a, _ := cache.TryGet("a")
	fmt.Printf("b evicted: %t, a=%d\n", !hasB, a)
}
