package client

import (
	"encoding/json"
)

func (c *Client) Location(location string) (RespLocation, error) {
	url := baseURL + "/location-area/" + location

	// Check cache first
	entry, found := c.cache.Get(url)
	if found {
		locationResp := RespLocation{}
		err := json.Unmarshal(entry, &locationResp)
		if err != nil {
			return RespLocation{}, err
		}

		return locationResp, nil
	}

	locationResp, dat, err := Call(url, RespLocation{}, *c)

	if err != nil {
		return RespLocation{}, err
	}

	// Add to cache
	c.cache.Add(url, dat)

	return locationResp, nil
}
