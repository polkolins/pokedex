package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.client.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locations.Next
	cfg.prevLocationsUrl = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("you are on the first page")
	}

	locations, err := cfg.client.ListLocations(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locations.Next
	cfg.prevLocationsUrl = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
