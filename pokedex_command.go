package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(confg *config, args ...string) error {
	if len(confg.caughtPokemon) == 0 {
		return errors.New("no pokemon in Pokedex")
	}

	fmt.Println("Pokemon in Pokedex")
	for _, pokemon := range confg.caughtPokemon {
		fmt.Printf("Name: %s\n", pokemon.Name)
	}

	return nil
}
