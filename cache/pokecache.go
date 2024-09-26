package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Map map[string]cacheEntry
	mu  *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Map: make(map[string]cacheEntry),
		mu:  &sync.RWMutex{},
	}
	go c.ReapLoop(interval)
	return c
}
