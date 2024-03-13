package main

import "fmt"

func commandHelp(c *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range getCommands(c) {
		help(cmd)
	}
	fmt.Println()
	return nil
}

func help(cmd cliCommand) {
	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
}
