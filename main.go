package main

import (
  "fmt"
  "strings"
  "os"
  "bufio"
  "internal/pokeapi"
)

var commands map[string]cliCommand
var configs config
var pokeApiClient pokeapi.PokeApiClient

func init() {
  commands = map[string]cliCommand {
    "help": {
      name:         "help",
      description:  "Displays a help message",
      callback:     func(config *config, param string) error {
        return commandHelp(commands)
      },
    },
    "exit": {
      name:         "exit",
      description:  "Exit the Pokedex",
      callback:     commandExit,
    },
    "map": {
      name:         "map",
      description:  "Display the names of the next 20 location areas in the Pokemon world",
      callback:     commandMap,
    },
    "mapb": {
      name:         "mapb",
      description:  "Display the names of the previous 20 location areas in the Pokemon world",
      callback:     commandMapb,
    },
    "explore": {
      name:         "explore",
      description:  "Display names of Pokemon in given area",
      callback:     commandExplore,
    },
    "catch": {
      name:         "catch",
      description:  "Try to catch a Pokemon with the given name",
      callback:     commandCatch,
    },
    "inspect": {
      name:         "inspect",
      description:  "Try to inspect a Pokemon with the given name in your Pokedex",
      callback:     commandInspect,
    },
    "pokedex": {
      name:         "pokedex",
      description:  "Show the names of the pokemon in your Pokedex",
      callback:     commandPokedex,
    },
  }

  configs = config{}

  pokeApiClient = pokeapi.NewClient()
}

func main() {
  inputScanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Pokedex > ")
    for inputScanner.Scan() {
      inputString := inputScanner.Text()
      inputs := cleanInput(inputString)

      if len(inputs) == 1 {
        command, ok := commands[inputs[0]]
        if !ok {
          fmt.Println("Unknown command")
          break
        }

        err := command.callback(&configs, "")
        if err != nil {
          fmt.Printf("Error when calling the command %s: %v\n", inputs[0], err)
        }

      } else if len(inputs) > 1 {
        command, ok := commands[inputs[0]]
        param := inputs[1]
        if !ok {
          fmt.Println("Unknown command")
          break
        }

        err := command.callback(&configs, param)
        if err != nil {
          fmt.Printf("Error when calling the command %s: %v\n", inputs[0], err)
        }

      } else {
        fmt.Println("No argument given, please provide at least one argument")
      }

      break
    }
  }
}

func cleanInput(text string) []string {
  fields := strings.Fields(strings.ToLower(text))
  return fields
}
