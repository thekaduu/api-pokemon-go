package pokemons

import (
	"database/sql"

	database "github.com/thekaduu/api-pokemon-go/packages/config"
)

type Pokemon struct {
	Id          int
	Name        string
	Type1       string
	Type2       *string
	Hp          int
	Attack      int
	Defense     int
	Speed       int
	Special     int
	Git         string
	Image       string
	Description string
}

func buildPokemons(rows *sql.Rows) ([]Pokemon, error) {
	var pokemonList []Pokemon

	for rows.Next() {
		var pokemon Pokemon
		err := rows.Scan(
			&pokemon.Id,
			&pokemon.Name,
			&pokemon.Type1,
			&pokemon.Type2,
			&pokemon.Hp,
			&pokemon.Attack,
			&pokemon.Defense,
			&pokemon.Speed,
			&pokemon.Special,
			&pokemon.Git,
			&pokemon.Image,
			&pokemon.Description)

		if err != nil {
			return nil, err
		}

		pokemonList = append(pokemonList, pokemon)
	}

	return pokemonList, nil
}

func GetPagedPokemons(page int) ([]Pokemon, error) {
	rows := database.Query("SELECT * FROM pokemons LIMIT 10 OFFSET $1", page*10)

	return buildPokemons(rows)
}

func All() ([]Pokemon, error) {
	rows := database.Query("SELECT * FROM pokemons")

	return buildPokemons(rows)
}

func Find(id int) (*Pokemon, error) {
	var pokemon Pokemon
	row := database.QueryRow("SELECT * FROM pokemons WHERE id = $1 LIMIT 1", id)

	err := row.Scan(
		&pokemon.Id,
		&pokemon.Name,
		&pokemon.Type1,
		&pokemon.Type2,
		&pokemon.Hp,
		&pokemon.Attack,
		&pokemon.Defense,
		&pokemon.Speed,
		&pokemon.Special,
		&pokemon.Git,
		&pokemon.Image,
		&pokemon.Description,
	)

	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}
