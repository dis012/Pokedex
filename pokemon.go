package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonData struct {
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
}

func (c *client) ListPokemon(api string) (Pokemon, error) {
	url := API_URL + api
	pokemon := Pokemon{}
	if val, ok := c.pokemonCache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshalling cached response: %v", err)
		}
		fmt.Println("Using cached response")
		c.currentPokemon = pokemon
		return pokemon, nil
	}

	res, err := c.http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error getting pokemons: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("error getting pokemons: %v", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	c.pokemonCache.Add(url, body)
	c.currentPokemon = pokemon
	return pokemon, nil
}

func (c *client) CatchPokemon(pokemon string) (PokemonData, error) {
	url := API_URL_POKEMON + pokemon
	coughtPokemon := PokemonData{}

	res, err := http.Get(url)
	if err != nil {
		return PokemonData{}, fmt.Errorf("error getting pokemon data: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return PokemonData{}, fmt.Errorf("error getting pokemon data: %v", res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &coughtPokemon)

	if err != nil {
		return PokemonData{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	base_experience := coughtPokemon.BaseExperience

	ifCaught := attemptCatch(base_experience)

	if !ifCaught {
		//fmt.Println("Pokemon got away")
		return PokemonData{}, fmt.Errorf("pokemon got away")
	}

	c.myCoughtPokemons.Add(pokemon, body)

	fmt.Println("Pokemon caught")

	return coughtPokemon, nil
}
