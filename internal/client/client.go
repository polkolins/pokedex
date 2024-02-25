package client

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/polkolins/pokedexcli/internal/cache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func NewClient(interval time.Duration) Client {
	return Client{
		cache:      cache.NewCache(interval),
		httpClient: http.Client{},
	}
}

func Call[Type any](url string, t Type, c Client) (Type, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return t, nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return t, nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, nil, err
	}

	tResp := t
	err = json.Unmarshal(dat, &resp)
	if err != nil {
		return t, nil, err
	}

	return tResp, dat, nil
}
