package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	all := cfg.deck.GetAll()

	if len(all) == 0 {
		fmt.Println("no pokemons in deck yet")
		return nil
	}

	for _, pokemon := range all {
		fmt.Printf("- %v\n", pokemon)
	}

	return nil
}
