package main

import (
	"time"

	"github.com/polkolins/pokedexcli/internal/client"
)

type config struct {
	client           client.Client
	deck             Deck
	nextLocationsUrl *string
	prevLocationsUrl *string
}

func main() {
	apiClient := client.NewClient(time.Second * 10)
	deck := NewDeck()

	cfg := &config{
		client: apiClient,
		deck:   deck,
	}

	startRepl(cfg)
}
