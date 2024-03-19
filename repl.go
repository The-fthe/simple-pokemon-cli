package main

import (
	"bufio"
	"fmt"
	"github.com/The-fthe/pokedex/internal/pokeapi"
	"os"
	"strings"
)

type Config struct {
	pokeapiClient   *pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	trainer         Trainer
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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, ok := getCommands()[cmdName]
		if ok {
			err := cmd.callback(c, args...)
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
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
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
		"explore": {
			name:        "explore <locaiton map>",
			description: "explore the location map",
			callback:    commandMapExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "try to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <catched pokemon name>",
			description: "try inspect catched pokemon",
			callback:    commandInspect,
		},
	}
}
