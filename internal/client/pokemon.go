package client

import (
	"encoding/json"
)

func (c *Client) Pokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	// Check cache first
	entry, found := c.cache.Get(url)
	if found {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(entry, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	pokemonResp, dat, err := Call(url, Pokemon{}, *c)

	if err != nil {
		return Pokemon{}, err
	}

	// Add to cache
	c.cache.Add(url, dat)

	return pokemonResp, nil
}
