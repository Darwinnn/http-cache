package main

import (
	"sync"
	"time"
)

// Cache represents the in-mem cache
type Cache struct {
	Map sync.Map
}

// CacheElement is the value stored in cache
type CacheElement struct {
	TTL     time.Time
	Content []byte
	Data    []byte
}

// Get returns CacheElement on key and bool wether the value has been found or not
// expired and notfound values return false
func (c *Cache) Get(key interface{}) (*CacheElement, bool) {
	if value, ok := c.Map.Load(key); ok {
		if value.(*CacheElement).TTL.Sub(time.Now()) <= 0 {
			c.Map.Delete(key)
			return nil, false
		}
		return value.(*CacheElement), true
	}
	return nil, false
}

// Set sets the value in cache
func (c *Cache) Set(key interface{}, value *CacheElement) {
	c.Map.Store(key, value)
}

// Del deletes data from cache by key
func (c *Cache) Del(key interface{}) {
	c.Map.Delete(key)
}

// DelExpired deletes all expired items from cache
func (c *Cache) DelExpired() {
	// TODO
}
