package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon to catch\n")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return errors.New("pokemon name is not valid\n")
	}
	randomCatchChance := rand.Intn(pokemon.BaseExperience)
	//fmt.Printf("base exp: %d , catch Chance: %d \n", pokemon.BaseExperience, randomCatchChance)

	if randomCatchChance > 40 {
		return fmt.Errorf("%s escaped!\n", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Printf("you may now inspect it with the inspect command.\n")
	c.trainer.AddToPokedex(pokemon)

	return nil
}
