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
	callback    func()
}

func startRepl() {
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

		availableCommands := getCommands()

		cliCommand, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command, command help is useful for you!")
			continue
		}

		cliCommand.callback()
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
	}
}
