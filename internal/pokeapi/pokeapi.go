package pokeapi

import (
	"encoding/json"
	"fmt"
	"internal/pokeapi/pokecache"
	"math/rand"
	"net/http"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var randomizer = rand.New(source)
var pokedex = make(map[string]Pokemon)

func NewClient() PokeApiClient {
	return PokeApiClient{
		httpClient: http.Client{
			Timeout: 5 * time.Second,
		},
		baseUrl: BASE_URL,
		cache:   pokecache.NewCache(5 * time.Second),
	}
}

func (pokeApiClient *PokeApiClient) GetLocationsFrom(pageUrl string) (LocationResponseBody, error) {
	url := pokeApiClient.baseUrl + "/location-area"
	locations := LocationResponseBody{}

	if pageUrl != "" {
		url = pageUrl
	}

	if item, ok := pokeApiClient.cache.Get(url); ok {
		fmt.Println("fetching from cache")
		err := json.Unmarshal(item, &locations)
		if err != nil {
			return locations, err
		}

		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()

	data, err := unmarshalJsonBodyIntoGivenStruct(res, &locations)
	if err != nil {
		return locations, err
	}

	pokeApiClient.cache.Add(url, data)

	return locations, err
}

func (pokeApiClient *PokeApiClient) GetPokemonFromLocation(locationName string) (ExplorationResponseBody, error) {
	url := pokeApiClient.baseUrl + "/location-area/" + locationName + "/"
	exploration := ExplorationResponseBody{}

	if item, ok := pokeApiClient.cache.Get(url); ok {
		fmt.Println("fetching from cache")
		err := json.Unmarshal(item, &exploration)
		if err != nil {
			return exploration, err
		}

		return exploration, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return exploration, err
	}
	defer res.Body.Close()

	data, err := unmarshalJsonBodyIntoGivenStruct(res, &exploration)
	if err != nil {
		return exploration, err
	}

	pokeApiClient.cache.Add(url, data)

	return exploration, nil
}

func (pokeApiClient *PokeApiClient) CatchPokemon(pokemonName string) (bool, error) {
	url := pokeApiClient.baseUrl + "/pokemon/" + pokemonName + "/"
	pokemon := Pokemon{}

	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	_, err = unmarshalJsonBodyIntoGivenStruct(res, &pokemon)
	if err != nil {
		return false, err
	}

	randomChance := randomizer.Intn(300)

	if randomChance < pokemon.BaseExperience {
		pokedex[pokemonName] = pokemon
		return true, nil
	}

	return false, nil
}

func Inspect(pokemonName string) {
	pokemon, ok := pokedex[pokemonName]
	if !ok {
		fmt.Printf("You have not caught %s!\n", pokemonName)
		return
	}
	pokemon.printDetails()
}
