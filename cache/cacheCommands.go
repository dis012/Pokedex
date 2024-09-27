package cache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Map[key] = cacheEntry{
		createdAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.Map[key]
	if !ok {
		return nil, false
	}
	return entry.Val, true
}

// if time.Duration stored in interval has passed, remove all entries older than that
func (c *Cache) ReapLoop(interval time.Duration) {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.Map {
		if now.Sub(entry.createdAt) > interval {
			delete(c.Map, key)
		}
	}
}

func (c *PokemonCache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Map[key] = cacheEntry{
		createdAt: time.Now(),
		Val:       val,
	}
}

func (c *PokemonCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.Map[key]
	if !ok {
		return nil, false
	}
	return entry.Val, true
}

func (c *PokemonCache) ReapLoop(interval time.Duration) {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.Map {
		if now.Sub(entry.createdAt) > interval {
			delete(c.Map, key)
		}
	}
}
