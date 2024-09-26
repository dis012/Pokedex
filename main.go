package main

import "time"

func main() {
	pokedexClient := newClient(5*time.Second, 5*time.Minute)
	cnf := Config{
		PokeClient: pokedexClient,
		NextPage:   nil,
		PrevPage:   nil,
	}
	pokedexClient.startCLI(&cnf)
}
