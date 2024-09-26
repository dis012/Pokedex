package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// map all available commands to their respective functions
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "Help",
			Description: "Displays help information",
			Command:     helpCommand,
		},
		"exit": {
			Name:        "Exit",
			Description: "Exits the program",
			Command:     exitCommand,
		},
		"map": {
			Name:        "Map",
			Description: "Displays a list of next locations",
			Command:     mapCommand,
		},
		"mapb": {
			Name:        "Mapb",
			Description: "Displays a list of previous locations",
			Command:     mapCommandBack,
		},
		"explore": {
			Name:        "Explore",
			Description: "Explore a location and find pokemons in that location",
			Command:     exploreCommand,
		},
	}
}

func (c *client) startCLI(cnf *Config) {
	commands := getCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		command := strings.ToLower(reader.Text())
		words := strings.Fields(command)
		if len(words) == 0 {
			continue
		}

		word := words[0]

		if word == "explore" && len(words) > 1 {
			location := words[1]
			cnf.Location = location
		}
		cliCommand, ok := commands[word]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := cliCommand.Command(cnf)
		if err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}
	}
}
