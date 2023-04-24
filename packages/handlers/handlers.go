package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pokemons "github.com/thekaduu/api-pokemon-go/packages/models"
	pokemonrepository "github.com/thekaduu/api-pokemon-go/packages/repositories"
)

func pagedPokemons(page int, c gin.Context) {
	pokemonList, err := pokemons.GetPagedPokemons(page)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": []pokemons.Pokemon{},
			"total":  0,
			"pages":  0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": pokemonList,
			"total":  len(pokemonList),
			"pages":  len(pokemonList) / 10,
		})
	}
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Recurso não encontrado.",
		"status":  http.StatusNotFound,
	})
}

func ListPokemons(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	var pokemonList []pokemons.Pokemon

	if err == nil {
		pagedPokemons(page, *c)
	} else {
		pokemonList, err = pokemonrepository.All()

		if pokemonList == nil {
			println(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, gin.H{
			"result": pokemonList,
			"total":  len(pokemonList),
			"pages":  len(pokemonList) / 10,
		})
	}
}

func ShowPokemon(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	pokemon, err := pokemonrepository.Find(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": nil,
			"status": http.StatusNotFound,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": pokemon,
	})
}
