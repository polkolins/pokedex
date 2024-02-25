package client

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache first
	entry, found := c.cache.Get(url)
	if found {
		locationsResp := RespLocations{}
		err := json.Unmarshal(entry, &locationsResp)
		if err != nil {
			return RespLocations{}, err
		}

		return locationsResp, nil
	}

	locationsResp, dat, err := Call(url, RespLocations{}, *c)

	if err != nil {
		return RespLocations{}, err
	}

	// Add to cache
	c.cache.Add(url, dat)

	return locationsResp, nil
}
