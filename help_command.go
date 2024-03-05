package main

import "fmt"

func callbackHelp() {
	fmt.Println("Welcome to Pokedex CLI tool!\n")
	fmt.Println("Here are some commands available for you:")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %v: %v\n", cmd.name, cmd.description)
	}

	fmt.Println()
}
