package chapter16

import "fmt"

type lruNode struct {
	key   string
	value int
	prev  *lruNode
	next  *lruNode
}

type LruCache struct {
	capacity int
	table    map[string]*lruNode
	head     *lruNode
	tail     *lruNode
}

func NewLruCache(capacity int) *LruCache {
	if capacity <= 0 {
		panic("Capacity must be positive.")
	}
	return &LruCache{capacity: capacity, table: make(map[string]*lruNode)}
}

func (c *LruCache) Count() int {
	return len(c.table)
}

func (c *LruCache) Get(key string) int {
	node, ok := c.table[key]
	if !ok {
		panic("key not found")
	}
	c.moveToLast(node)
	return node.value
}

func (c *LruCache) Put(key string, value int) {
	if existing, ok := c.table[key]; ok {
		existing.value = value
		c.moveToLast(existing)
		return
	}
	if len(c.table) >= c.capacity {
		oldest := c.head
		c.remove(oldest)
		delete(c.table, oldest.key)
	}
	node := &lruNode{key: key, value: value}
	c.append(node)
	c.table[key] = node
}

func (c *LruCache) moveToLast(node *lruNode) {
	c.remove(node)
	c.append(node)
}

func (c *LruCache) remove(node *lruNode) {
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

func (c *LruCache) append(node *lruNode) {
	if c.tail == nil {
		c.head = node
		c.tail = node
		return
	}
	c.tail.next = node
	node.prev = c.tail
	c.tail = node
}

func RunQ1625() {
	cache := NewLruCache(2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	fmt.Println(cache.Get("a"))
	cache.Put("c", 3)
	fmt.Printf("Count=%d\n", cache.Count())
}
