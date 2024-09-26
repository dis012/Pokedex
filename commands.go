package main

import (
	"fmt"
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
