package main

import (
	"errors"
	"fmt"
)

func callbackExplore(confg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location area provided")
	} else if len(args) > 1 {
		return errors.New("one location area is needed")
	}

	locationAreaName := args[0]

	locationArea, err := confg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
