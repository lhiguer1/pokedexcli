package main

import (
	"time"

	"github.com/lhiguer1/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeAPIClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeAPIClient: pokeapi.NewClient(time.Minute, time.Hour),
	}
	startRepl(&cfg)
}
