module github.com/wittyCode/pokedexcli

go 1.24.0

require internal/pokeapi v0.0.1

replace internal/pokeapi => ./internal/pokeapi/

require internal/pokeapi/pokecache v0.0.1

replace internal/pokeapi/pokecache => ./internal/pokeapi/pokecache/
