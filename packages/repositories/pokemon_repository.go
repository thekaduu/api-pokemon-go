package pokemonrepository

import (
	"encoding/json"

	redis "github.com/thekaduu/api-pokemon-go/packages/config"
	pokemons "github.com/thekaduu/api-pokemon-go/packages/models"
)

func All() ([]pokemons.Pokemon, error) {
	var pokemonsList []pokemons.Pokemon
	cachedPokemons, _ := redis.Get("all-pokemons")

	if cachedPokemons == "" {
		pokemonsList, _ = pokemons.All()
		pokemonJson, _ := json.Marshal(pokemonsList)

		redis.Set("all-pokemons", pokemonJson)
		return pokemonsList, nil
	} else {
		err := json.Unmarshal([]byte(cachedPokemons), &pokemonsList)
		if err != nil {
			return nil, err
		}
	}

	return pokemonsList, nil
}

func Find(id int) (*pokemons.Pokemon, error) {
	cachedKey := "pokemon-id-" + string(id)
	cachedPokemon, _ := redis.Get(cachedKey)
	var pokemon pokemons.Pokemon

	if cachedPokemon == "" {
		pokemon, _ := pokemons.Find(id)
		pokemonJson, _ := json.Marshal(pokemon)

		redis.Set(cachedKey, pokemonJson)
		return pokemon, nil
	}

	json.Unmarshal([]byte(cachedPokemon), &pokemon)

	if pokemon.Id == 0 {
		return nil, nil
	}

	return &pokemon, nil
}
