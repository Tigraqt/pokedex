package main

import (
	"time"

	"github.com/Tigraqt/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeClient          pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	config := config{
		pokeClient:    pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&config)
}
