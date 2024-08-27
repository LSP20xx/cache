package cache

import (
	"errors"
	"sync"
)

type Cache[K comparable, V any] struct {
	data       map[K]V
	mu         sync.Mutex
	concurrent bool
}

func NewCache[K comparable, V any](concurrent bool) *Cache[K, V] {
	return &Cache[K, V]{
		data:       make(map[K]V),
		concurrent: concurrent,
	}
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	if c.concurrent {
		c.mu.Lock()
		defer c.mu.Unlock()
	}

	value, exists := c.data[key]
	if !exists {
		var zeroValue V
		return zeroValue, errors.New("key not found in cache")
	}

	return value, nil
}

func (c *Cache[K, V]) Set(key K, value V) {
	if c.concurrent {
		c.mu.Lock()
		defer c.mu.Unlock()
	}

	c.data[key] = value
}

func (c *Cache[K, V]) Delete(key K) {
	if c.concurrent {
		c.mu.Lock()
		defer c.mu.Unlock()
	}

	delete(c.data, key)
}

func (c *Cache[K, V]) Clear() {
	if c.concurrent {
		c.mu.Lock()
		defer c.mu.Unlock()
	}

	c.data = make(map[K]V)
}
