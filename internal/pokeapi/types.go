package pokeapi

import (
	"internal/pokeapi/pokecache"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
}

type LocationResponseBody struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int             `json:"base_stat"`
	Effort   int             `json:"effort"`
	Stat     PokemonStatType `json:"stat"`
}

type PokemonStatType struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type PokemonTypeDef `json:"type"`
}

type PokemonTypeDef struct {
	Name string `json:"name"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type ExplorationResponseBody struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokeApiClient struct {
	httpClient http.Client
	baseUrl    string
	cache      pokecache.Cache
}
