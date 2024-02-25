package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("pokemon name missing")
	}

	name := args[0]
	pokemon, found := cfg.deck.Get(name)

	if !found {
		return errors.New("pokemon " + name + "is not found")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Heigh: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}
