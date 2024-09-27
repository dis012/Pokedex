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
	Val       []byte
}

type PokemonCache struct {
	Map map[string]cacheEntry
	mu  *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Map: make(map[string]cacheEntry),
		mu:  &sync.RWMutex{},
	}
	go c.ReapLoop(interval)
	return c
}

func NewPokemonCache(interval time.Duration) PokemonCache {
	c := PokemonCache{
		Map: make(map[string]cacheEntry),
		mu:  &sync.RWMutex{},
	}
	go c.ReapLoop(interval)
	return c
}
