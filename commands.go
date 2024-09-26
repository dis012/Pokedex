package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// help command which displays all available commands and ther functions
func helpCommand(cnf *Config) error {
	commands := getCommands() // map[string]cliCommand
	if commands == nil {
		return fmt.Errorf("no commands found")
	}
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

// exit command which exits the program
func exitCommand(cnf *Config) error {
	fmt.Println("Exiting program")
	os.Exit(0)
	return nil
}

// map command which displays a list of locations
func mapCommand(cnfg *Config) error {
	location := PokeMap{}
	if cnfg.NextPage == nil {
		location, err := cnfg.PokeClient.ListLocations(nil)
		if err != nil {
			return fmt.Errorf("error getting locations: %v", err)
		}
		for _, loc := range location.Results {
			fmt.Println(loc.Name)
		}
		cnfg.NextPage = &location.NextPage
		return nil
	}

	location, err := cnfg.PokeClient.ListLocations(cnfg.NextPage)
	if err != nil {
		return fmt.Errorf("error getting locations: %v", err)
	}
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	cnfg.NextPage = &location.NextPage
	cnfg.PrevPage = &location.PrevPage
	return nil
}

func mapCommandBack(cnf *Config) error {
	if cnf.PrevPage == nil {
		fmt.Println("No previous page")
		return nil
	}

	location, err := cnf.PokeClient.ListLocations(cnf.PrevPage)
	if err != nil {
		return fmt.Errorf("error getting locations: %v", err)
	}
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	cnf.NextPage = &location.NextPage
	cnf.PrevPage = &location.PrevPage
	return nil
}

func exploreCommand(cnf *Config) error {
	fmt.Println("Exploring ", cnf.Location)
	var pokemon Pokemon
	if cnf.Location == "" {
		fmt.Println("No location provided")
		return nil
	}

	// Check if the location exists in the cache
	err := cnf.checkLocation()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Get the pokemons in the location
	url := API_URL + cnf.Location
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting pokemons: %v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("error getting pokemons: %s", body)
	}
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %v", err)
	}

	for _, pok := range pokemon.PokemonEncounters {
		fmt.Println(pok.Pokemon.Name)
	}

	return nil
}
