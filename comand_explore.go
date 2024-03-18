package main

import (
	"errors"
	"fmt"
)

func commandMapExplore(c *Config) error {
	ExpRep, err := c.pokeapiClient.ExploreLocation(c.parameter)
	if err != nil {
		return err
	}
	if c.parameter == nil {
		return errors.New("parameter is empty")
	}

	fmt.Printf("Exploring %s...\n", *c.parameter)

	for _, encounter := range ExpRep.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}
	return nil

}
