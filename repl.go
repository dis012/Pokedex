package main

import (
	"fmt"
	"strings"
)

// cliCommand struct which holds the name, description and command function and
// is used to store all available commands
type cliCommand struct {
	Name        string
	Description string
	Command     func(cnf *config) error
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

func startCLI() {
	commands := getCommands()
	var cnf config

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

		err := cliCommand.Command(&cnf)
		if err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}
	}
}
