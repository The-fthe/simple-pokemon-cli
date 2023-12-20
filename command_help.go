package main

import "fmt"

func commandHelp() error{
  fmt.Println()
  fmt.Println("Welcome to the Pokedex!")
  fmt.Println("Usage:")
  for _,cmd := range getCommands(){
    help(cmd)
  }
  fmt.Println()
  return nil
}

func help(cmd cliCommand){
  fmt.Printf("%s: %s\n",cmd.name,cmd.description)
}
