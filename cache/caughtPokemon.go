package cache

import "sync"

type MyPokemonsCache struct {
	Map map[string]pokemonEntry
	mu  *sync.RWMutex
}

type pokemonEntry struct {
	Val []byte
}

func NewMyPokemonsCache() MyPokemonsCache {
	return MyPokemonsCache{
		Map: make(map[string]pokemonEntry),
		mu:  &sync.RWMutex{},
	}
}
