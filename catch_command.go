package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(confg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	} else if len(args) > 1 {
		return errors.New("one pokemon name is needed")
	}

	pokemonName := args[0]

	pokemon, err := confg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 30
	randNum := rand.Intn(pokemon.BaseExperience)

	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
