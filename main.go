package main

import (
	"time"

	"github.com/vnsonvo/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	confg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 10),
	}

	startRepl(&confg)
}
