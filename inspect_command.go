package main

import (
	"errors"
	"fmt"
)

func callbackInsepct(confg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	} else if len(args) > 1 {
		return errors.New("one pokemon name is needed")
	}

	pokemonName := args[0]
	pokemon, ok := confg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, tp := range pokemon.Types {
		fmt.Printf(" - %s\n", tp.Type.Name)
	}

	return nil
}
