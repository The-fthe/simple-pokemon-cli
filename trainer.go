package main

import (
	"fmt"

	"github.com/The-fthe/pokedex/internal/pokeapi"
)

type Trainer struct {
	Pokedex map[string]pokeapi.Pokemon `json:"pokemon"`
}

func (t *Trainer) AddToPokedex(p pokeapi.Pokemon) {
	if t.Pokedex == nil {
		t.NewPokedex()
	}
	if _, ok := t.Pokedex[p.Name]; ok {
		fmt.Printf("%s is recorded\n", p.Name)
		return
	}
	t.Pokedex[p.Name] = p
}

func (t *Trainer) NewPokedex() {
	t.Pokedex = make(map[string]pokeapi.Pokemon)
}
