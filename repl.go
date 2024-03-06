package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(confg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		cleanedInput := cleanInput(text)
		commandName := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		availableCommands := getCommands()

		cliCommand, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command, command help is useful for you!")
			continue
		}

		err := cliCommand.callback(confg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	loweredInput := strings.ToLower(str)
	words := strings.Fields(loweredInput)
	return words
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List some previous location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Attempt to catch a pokemon and add it to pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "View information about a caught pokemon",
			callback:    callbackInsepct,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the pokemon in your pokedex",
			callback:    callbackPokedex,
		},
	}
}
