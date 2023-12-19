package main

import(
  "fmt"
  "os"
  "strings"
  "bufio"
  "errors"
)


func main() {
  scanner := bufio.NewScanner(os.Stdin)
  cmdMap := InitCommand()
  fmt.Printf("Pokedex >")
  for scanner.Scan(){
    input := scanner.Text()
    if strings.ToLower(input) == "exit"{
      break
    }
    if strings.ToLower(input) == "help"{
      fmt.Println("Welcome to the Pokedex!")
      fmt.Println("Usage:\n")
      help(cmdMap["help"])
      help(cmdMap["exit"])
    }
    fmt.Printf("Pokedex >")
  }
  if err := scanner.Err();err != nil{
    fmt.Fprintln(os.Stderr,"Error reading from input:",err)

  }
}

func InitCommand() map[string]cliCommand{
  commandExit := func()error{return errors.New("Exit error")}

  commandHelp := func()error{return errors.New("Help error")}

  return map[string]cliCommand{
    "help":{
      name: "help",
      description: "Displays a help message",
      callback: commandHelp,
    },
    "exit":{
      name: "exit",
      description: "Exit the Pokedex",
      callback: commandExit,
    },
  }
}


func help(cmd cliCommand){
  fmt.Printf("%s: %s\n",cmd.name,cmd.description)
}





type cliCommand struct{
  name string
  description string
  callback func()error
}

