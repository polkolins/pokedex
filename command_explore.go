package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("location missing")
	}

	name := args[0]
	fmt.Printf("Exploring %v...\n", name)
	locations, err := cfg.client.Location(name)
	if err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")
	for _, enc := range locations.PokemonEncounters {
		fmt.Printf("- %s\n", enc.Pokemon.Name)
	}
	return nil
}
