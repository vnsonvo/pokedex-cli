package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		fmt.Println("Text:", text)
	}
}

func cleanInput(str string) []string {
	loweredInput := strings.ToLower(str)
	words := strings.Fields(loweredInput)
	return words
}
