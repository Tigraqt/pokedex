package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const tresHold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	fmt.Println(pokemon.BaseExperience, randNum, tresHold)

	if randNum > tresHold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
