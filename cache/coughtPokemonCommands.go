package cache

func (pc *MyPokemonsCache) Add(key string, val []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.Map[key] = pokemonEntry{
		Val: val,
	}
}

func (pc *MyPokemonsCache) GetAll() map[string]pokemonEntry {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.Map
}

func (pc *MyPokemonsCache) GetPokemonStats(pokemon string) []byte {
	pc.mu.RLock()
	defer pc.mu.Unlock()
	key := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	entry, ok := pc.Map[key]
	if !ok {
		return nil
	}
	return entry.Val
}
