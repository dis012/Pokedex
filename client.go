package main

import (
	"fmt"
	"net/http"
	"pokedexcli/cache"
	"time"
)

type Config struct {
	PokeClient client
	Location   string  // Location which the user wants to search for pokemons
	Pokemon    string  // Pokemon which the user wants to search for
	NextPage   *string // Next page of the API
	PrevPage   *string // Previous page of the API
}

// cliCommand struct which holds the name, description and command function and
// is used to store all available commands
type cliCommand struct {
	Name        string
	Description string
	Command     func(cnf *Config) error
}

// Client is used to interact with the pokedex API
// and cache the responses
type client struct {
	cache            cache.Cache
	pokemonCache     cache.PokemonCache
	currentArea      PokeMap
	currentPokemon   Pokemon
	myCoughtPokemons cache.MyPokemonsCache
	http             http.Client
}

func newClient(timeout, interval time.Duration) client {
	return client{
		cache:            cache.NewCache(interval),
		pokemonCache:     cache.NewPokemonCache(interval),
		myCoughtPokemons: cache.NewMyPokemonsCache(),
		http: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Config) checkLocation() error {
	for _, loc := range c.PokeClient.currentArea.Results {
		if loc.Name == c.Location {
			return nil
		}
	}
	return fmt.Errorf("location not found")
}

func (c *Config) checkPokemon() error {
	for _, poke := range c.PokeClient.currentPokemon.PokemonEncounters {
		if poke.Pokemon.Name == c.Pokemon {
			return nil
		}
	}

	return fmt.Errorf("pokemon not found in the current location")
}
