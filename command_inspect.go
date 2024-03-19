package main

import (
	"errors"
)

func commandInspect(c *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon to catch\n")
	}

	pokemonName := args[0]
	err := c.trainer.InspectPokemon(pokemonName)
	if err != nil {
		return err
	}

	return nil
}
