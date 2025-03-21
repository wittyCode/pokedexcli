package main

import (
	"errors"
	"fmt"
	"internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, param string) error
}

type config struct {
	Next     string
	Previous string
}

func commandExit(config *config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commandMap map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println("")

	return nil
}

func commandMap(config *config, param string) error {
	data, err := pokeApiClient.GetLocationsFrom(config.Next)
	if err != nil {
		return err
	}

	url := data.Next
	if url == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	printLocations(data.Results)
	updateConfig(config, data.Next, data.Previous)
	return nil
}

func commandMapb(config *config, param string) error {
	url := config.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	data, err := pokeApiClient.GetLocationsFrom(url)
	if err != nil {
		return err
	}

	printLocations(data.Results)
	updateConfig(config, data.Next, data.Previous)
	return nil
}

func commandExplore(config *config, param string) error {
	if param == "" {
		return errors.New("you need to give a location name")
	}

	data, err := pokeApiClient.GetPokemonFromLocation(param)

	if err != nil {
		return err
	}

	printPokemon(data.PokemonEncounters)
	return nil
}

func commandCatch(config *config, param string) error {
	if param == "" {
		return errors.New("you need to give a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", param)
	success, err := pokeApiClient.CatchPokemon(param)

	if err != nil {
		return err
	}

	if success {
		fmt.Printf("%s was caught!\n", param)
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", param)
	}

	return nil
}

func commandInspect(config *config, param string) error {
	if param == "" {
		return errors.New("you need to give a pokemon name")
	}

	pokeapi.Inspect(param)
	return nil
}

func commandPokedex(config *config, param string) error {
	pokeapi.PrintPokedex()
	return nil
}

func printLocations(locations []pokeapi.LocationArea) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func printPokemon(pokemonEncounters []pokeapi.PokemonEncounter) {
	for _, pok := range pokemonEncounters {
		fmt.Println(pok.Pokemon.Name)
	}
}

func updateConfig(configPtr *config, next, prev string) {
	configPtr.Next = next
	configPtr.Previous = prev
}
