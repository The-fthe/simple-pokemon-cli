package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/The-fthe/pokedex/internal/pokeapi"
)

func commandMapForward(c *Config) error {
	var locationsResp pokeapi.RespLocations
	var err error

	if c.nextLocationURL != nil {
		dat, ok := c.pokeCache.Get(*c.nextLocationURL)
		if !ok {
			locationsResp, err = c.pokeapiClient.ListLocation(c.nextLocationURL)
			if err != nil {
				return err
			}
		} else {
			err := json.Unmarshal(dat, &locationsResp)
			if err != nil {
				return err
			}
		}
	} else {
		locationsResp, err = c.pokeapiClient.ListLocation(c.nextLocationURL)
		if err != nil {
			return err
		}
	}

	c.nextLocationURL = locationsResp.Next
	c.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapPrevious(c *Config) error {
	if c.prevLocationURL == nil {
		return errors.New("you are in first page")
	}

	var locationsResp pokeapi.RespLocations
	var err error

	dat, ok := c.pokeCache.Get(*c.prevLocationURL)
	ok = false
	if !ok {
		locationsResp, err = c.pokeapiClient.ListLocation(c.prevLocationURL)
		if err != nil {
			return err
		}
	} else {
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return err
		}
	}

	c.nextLocationURL = locationsResp.Next
	c.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
