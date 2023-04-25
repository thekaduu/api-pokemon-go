package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	pokemons "github.com/thekaduu/api-pokemon-go/packages/models"
	pokemonrepository "github.com/thekaduu/api-pokemon-go/packages/repositories"
)

func pagedPokemons(page int, c fiber.Ctx) error {
	pokemonList, err := pokemons.GetPagedPokemons(page)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": []pokemons.Pokemon{},
			"total":  0,
			"pages":  0,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": pokemonList,
		"total":  len(pokemonList),
		"pages":  len(pokemonList) / 10,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Recurso não encontrado.",
		"status":  http.StatusNotFound,
	})
}

func ListPokemons(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", ""))
	var pokemonList []pokemons.Pokemon

	if err == nil {
		return pagedPokemons(page, *c)
	}

	pokemonList, err = pokemonrepository.All()

	if pokemonList == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"rersult": nil,
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": pokemonList,
		"total":  len(pokemonList),
		"pages":  len(pokemonList) / 10,
	})
}

func ShowPokemon(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	pokemon, err := pokemonrepository.Find(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"result": nil,
			"status": fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": pokemon})
}
