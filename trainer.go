package main

import (
	"errors"
	"fmt"

	"github.com/The-fthe/pokedex/internal/pokeapi"
)

type Trainer struct {
	CaughtPokemon map[string]pokeapi.Pokemon `json:"pokemon"`
}

func (t *Trainer) AddToPokedex(p pokeapi.Pokemon) {
	if t.CaughtPokemon == nil {
		t.NewPokedex()
	}
	if _, ok := t.CaughtPokemon[p.Name]; ok {
		fmt.Printf("%s is recorded\n", p.Name)
		return
	}
	t.CaughtPokemon[p.Name] = p
}

func (t *Trainer) NewPokedex() {
	t.CaughtPokemon = make(map[string]pokeapi.Pokemon)
}

func (t *Trainer) InspectPokemon(pokemonName string) error {
	if pokemon, ok := t.CaughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s \n", pokemon.Name)
		fmt.Printf("Height: %s \n", pokemon.Name)
		fmt.Printf("Weight: %s \n", pokemon.Name)
		fmt.Printf("Stats: \n")
		fmt.Printf("  -hp: %d \n", GetBaseStat(pokemon, "hp"))
		fmt.Printf("  -attack: %d \n", GetBaseStat(pokemon, "attack"))
		fmt.Printf("  -defense: %d \n", GetBaseStat(pokemon, "defense"))
		fmt.Printf("  -special-attack: %d \n", GetBaseStat(pokemon, "special-attack"))
		fmt.Printf("  -special-defense: %d \n", GetBaseStat(pokemon, "special-defense"))
		fmt.Printf("  -speed: %d \n", GetBaseStat(pokemon, "speed"))
		fmt.Printf("Types: \n")
		PrintType(pokemon)
		return nil
	}
	return errors.New("you have not caught that pokemon")
}
func (t *Trainer) PrintContainPokemon() error {
	if t.CaughtPokemon == nil {
		return errors.New("Pokedex is empty")
	}
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range t.CaughtPokemon {
		fmt.Printf(" -%s\n", pokemon.Name)
	}
	return nil
}

func GetBaseStat(p pokeapi.Pokemon, statName string) int {
	for _, stat := range p.Stats {
		if stat.Stat.Name == statName {
			return stat.BaseStat
		}
	}
	fmt.Printf("%s is not available", statName)
	return 0
}

func PrintType(p pokeapi.Pokemon) {
	for _, t := range p.Types {
		fmt.Printf("  - %s \n", t.Type.Name)
	}
}
