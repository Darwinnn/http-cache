package main

import (
	"sync"
	"time"
)

// Cache represents the in-mem cache
type Cache struct {
	Mutex *sync.Mutex
	Map   map[string]CacheElement
}

// CacheElement is the value stored in cache
type CacheElement struct {
	TTL     time.Time
	Content string
	Data    []byte
}

// Get returns CacheElement on key and bool wether the value has been found or not
// expired and notfound values return false
func (c *Cache) Get(key string) (CacheElement, bool) {
	if value, ok := c.Map[key]; ok {
		if value.TTL.Sub(time.Now()) <= 0 {
			delete(c.Map, key)
			return CacheElement{}, false
		}
		return value, true
	}
	return CacheElement{}, false
}

// Set sets the value in cache
func (c *Cache) Set(key string, value CacheElement) {
	c.Mutex.Lock()
	c.Map[key] = value
	c.Mutex.Unlock()
}

// Del deletes data from cache by key
func (c *Cache) Del(key string) {
	delete(c.Map, key)
}

// DelExpired deletes all expired items from cache
func (c *Cache) DelExpired() {
	for k, v := range c.Map {
		if v.TTL.Sub(time.Now()) <= 0 {
			delete(c.Map, k)
		}
	}
}
