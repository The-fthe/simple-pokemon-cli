package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/The-fthe/pokedex/internal/pokeapi"
	"github.com/The-fthe/pokedex/internal/pokecache"
)

type Config struct {
	pokeapiClient   pokeapi.Client
	pokeCache       *pokecache.Cache
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(c *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		cmd, ok := getCommands(c)[cmdName]
		if ok {
			err := cmd.callback(c)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands(c *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "display next 20 locations",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "map back",
			description: "display previous 20 locations",
			callback:    commandMapPrevious,
		},
	}
}
