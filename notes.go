package main

/*
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// help command which displays all available commands and ther functions
func helpCommand(cnf *config) error {
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
func exitCommand(cnf *config) error {
	fmt.Println("Exiting program")
	os.Exit(0)
	return nil
}

func getFirstPage(cnf *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return fmt.Errorf("error getting first page: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("error getting first page: %s", body)
	}
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, cnf)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %v", err)
	}

	return nil
}

// map command which displays a list of locations
func mapCommand(cnfg *config) error {
	if cnfg.NextPage == "" {
		err := getFirstPage(cnfg)
		if err != nil {
			return fmt.Errorf("error getting first page: %v", err)
		}

		fmt.Println("Locations:")
		if cnfg.Results == nil {
			return fmt.Errorf("no results found")
		}
		for _, loc := range cnfg.Results {
			fmt.Println(loc.Name)
		}
		return nil
	} else {
		res, err := http.Get(cnfg.NextPage)
		if err != nil {
			return fmt.Errorf("error getting next page: %v", err)
		}

		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("error getting next page: %s", body)
		}
		if err != nil {
			return fmt.Errorf("error reading response body: %v", err)
		}

		err = json.Unmarshal(body, cnfg)
		if err != nil {
			return fmt.Errorf("error unmarshalling response: %v", err)
		}

		fmt.Println("Locations:")
		for _, loc := range cnfg.Results {
			fmt.Println(loc.Name)
		}

		return nil
	}
}

func mapCommandBack(cnf *config) error {
	if cnf.PrevPage == "" {
		return fmt.Errorf("no previous page found")
	}
	res, err := http.Get(cnf.PrevPage)
	if err != nil {
		return fmt.Errorf("error getting previous page: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("error getting previous page: %s", body)
	}
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, cnf)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %v", err)
	}

	fmt.Println("Locations:")
	for _, loc := range cnf.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

*/
