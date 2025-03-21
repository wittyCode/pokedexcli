package main

import (
	"bufio"
	"fmt"
	"internal/pokeapi"
	"os"
	"strings"
)

var commands map[string]cliCommand
var configs config
var pokeApiClient pokeapi.PokeApiClient

func init() {
	commands = initMap()

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
