package main

import "github.com/vnsonvo/pokedex-cli/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	confg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&confg)
}
