package main

import (
	"log"

	"github.com/joho/godotenv"
	handlers "github.com/thekaduu/api-pokemon-go/packages/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	r := gin.Default()
	r.NoRoute(handlers.NotFound)

	r.GET("/pokemons", handlers.ListPokemons)
	r.GET("pokemons/:id", handlers.ShowPokemon)

	r.Run()
}
