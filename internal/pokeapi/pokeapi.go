package pokeapi

import (
  "fmt"
  "net/http"
  "io"
  "encoding/json"
  "time"
  "internal/pokeapi/pokecache"
  "math/rand"
)

type LocationArea struct {
  Name string `json:"name"`
}

type LocationResponseBody struct {
  Count    int          `json:"count"`
  Next     string       `json:"next"`
  Previous string          `json:"previous"`
  Results  []LocationArea `json:"results"`
}

type Pokemon struct {
  Name            string          `json:"name"`
  BaseExperience  int             `json:"base_experience"`
  Height          int             `json:"height"`
  Weight          int             `json:"weight"`
  Stats           []PokemonStat   `json:"stats"`
  Types           []PokemonType   `json:"types"`
}

type PokemonStat struct {
  BaseStat  int             `json:"base_stat"`
  Effort    int             `json:"effort"`
  Stat      PokemonStatType `json:"stat"`
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
  baseUrl string
  cache pokecache.Cache
}

var source = rand.NewSource(time.Now().UnixNano())
var randomizer = rand.New(source)
var Pokedex = make(map[string]Pokemon)

func NewClient() PokeApiClient {
  return PokeApiClient {
    httpClient: http.Client{
      Timeout: 5 * time.Second,
    },
    baseUrl: BASE_URL,
    cache: pokecache.NewCache(5 * time.Second),
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

  data, err := io.ReadAll(res.Body)
  if err != nil {
    return locations, err
  }

  pokeApiClient.cache.Add(url, data)
  err = json.Unmarshal(data, &locations)
  if err != nil {
    return locations, err
  }

  return locations, nil
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

  data, err := io.ReadAll(res.Body)
  if err != nil {
    return exploration, err
  }

  pokeApiClient.cache.Add(url, data)
  err = json.Unmarshal(data, &exploration)
  if err != nil {
    return exploration, err
  }

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

  data, err := io.ReadAll(res.Body)
  if err != nil {
    return false, err
  }

  err = json.Unmarshal(data, &pokemon)
  if err != nil {
    return false, err
  }

  randomChance := randomizer.Intn(300)

  if randomChance < pokemon.BaseExperience {
    Pokedex[pokemonName] = pokemon
    return true, nil
  }

  return false, nil
}

func Inspect(pokemonName string) {
  pokemon, ok := Pokedex[pokemonName]
  if !ok {
    fmt.Printf("You have not caught %s!\n", pokemonName)
    return
  }
  pokemon.printDetails()
}

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
  for key := range Pokedex {
    fmt.Printf("  - %s\n", key)
  }
}
