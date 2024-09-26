package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"pokedexcli/cache"
)

// cliCommand struct which holds the name, description and command function and
// is used to store all available commands
type cliCommand struct {
	Name        string
	Description string
	Command     func(cnf *Config) error
}

type client struct {
	cache cache.Cache
	http  http.Client
}

func newClient(timeout, interval time.Duration) client {
	return client{
		cache: cache.NewCache(interval),
		http: http.Client{
			Timeout: timeout,
		},
	}
}

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
	}
}

// getInput function which reads input from the user
func getInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func (c *client) startCLI(cnf *Config) {
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		command := strings.ToLower(getInput())
		if len(command) == 0 {
			continue
		}
		cliCommand, ok := commands[command]

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
