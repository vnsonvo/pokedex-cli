package main

import (
	"time"

	"github.com/vnsonvo/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	confg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 10),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&confg)
}
