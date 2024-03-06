package main

import "fmt"

func callbackHelp(confg *config) error {
	fmt.Println("Welcome to Pokedex CLI tool!")
	fmt.Println("Here are some commands available for you:")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %v: %v\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
