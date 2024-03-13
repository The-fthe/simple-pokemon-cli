package main

import "fmt"

func (c *Config) commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range c.getCommands() {
		help(cmd)
	}
	fmt.Println()
	return nil
}

func help(cmd cliCommand) {
	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
}
