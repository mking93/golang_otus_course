package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]cacheItem
}

type cacheItem struct {
	key   *ListItem
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]cacheItem),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	cItem, found := c.items[key]
	var cacheKey *ListItem

	if found {
		cacheKey = cItem.key
		c.queue.MoveToFront(cacheKey)
	} else {
		cacheKey = c.queue.PushFront(key)
	}

	if c.queue.Len() > c.capacity {
		lastQItem := c.queue.Back()

		delete(c.items, lastQItem.Value.(Key))
		c.queue.Remove(lastQItem)
	}

	c.items[key] = cacheItem{value: value, key: cacheKey}
	return found
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	cacheItem, found := c.items[key]
	if found {
		c.queue.MoveToFront(cacheItem.key)
	}

	return cacheItem.value, found
}

func (c *lruCache) Clear() {
	c.queue = NewList()

	for key := range c.items {
		delete(c.items, key)
	}
}
