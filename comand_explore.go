package main

import (
	"errors"
	"fmt"
)

func commandMapExplore(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	locationName := args[0]
	ExpRep, err := c.pokeapiClient.ExploreLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)

	for _, encounter := range ExpRep.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}
	return nil

}
