package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("pokemon name missing")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	pokemon, err := cfg.client.Pokemon(name)
	if err != nil {
		return err
	}

	exp := pokemon.BaseExperience
	catch := rand.Intn(exp) >= exp/2

	if !catch {
		fmt.Printf("%v escaped!\n", name)
		return nil
	}

	fmt.Printf("%v was caught!\n", name)
	cfg.deck.Add(pokemon.Name, pokemon)

	return nil
}
