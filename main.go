package main

import (
	"time"

	"github.com/The-fthe/pokedex/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
