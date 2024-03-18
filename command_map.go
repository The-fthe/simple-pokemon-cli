package main

import (
	"errors"
	"fmt"
)

func commandMapForward(c *Config) error {

	locationsResp, err := c.pokeapiClient.ListLocation(c.nextLocationURL)
	if err != nil {
		return err
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

	locationsResp, err := c.pokeapiClient.ListLocation(c.prevLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locationsResp.Next
	c.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
