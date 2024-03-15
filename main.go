package main

import (
	"github.com/The-fthe/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
