package main

import (
	"fmt"
	"time"

	"github.com/The-fthe/pokedex/internal/pokeapi"
	"github.com/The-fthe/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	if pokeCache == nil {
		fmt.Printf("pokeCache is nil!!")
		return
	}

	cfg := &Config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}

	startRepl(cfg)
}
