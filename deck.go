package main

import (
	"github.com/polkolins/pokedexcli/internal/client"
	"sync"
)

type Deck struct {
	deck map[string]client.Pokemon
	mux  *sync.Mutex
}

func NewDeck() Deck {
	deck := Deck{
		deck: make(map[string]client.Pokemon),
		mux:  &sync.Mutex{},
	}
	return deck
}

func (p *Deck) Add(key string, pokemon client.Pokemon) {
	p.deck[key] = pokemon
}

func (p *Deck) Get(key string) (client.Pokemon, bool) {
	pokemon, found := p.deck[key]
	return pokemon, found
}

func (p *Deck) GetAll() []string {
	var all []string

	for _, pokemon := range p.deck {
		all = append(all, pokemon.Name)
	}

	return all
}
