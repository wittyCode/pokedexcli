package pokeapi

import (
	"fmt"
)

func (pokemon *Pokemon) printDetails() {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		statValue := stat.BaseStat + stat.Effort
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, statValue)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
}

func PrintPokedex() {
	fmt.Println("Your Pokedex:")
	for key := range pokedex {
		fmt.Printf("  - %s\n", key)
	}
}
