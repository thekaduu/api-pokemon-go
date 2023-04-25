package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	handlers "github.com/thekaduu/api-pokemon-go/packages/handlers"
)

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}
}

func main() {
	loadEnvVariables()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Use(logger.New(logger.Config{
		Format: "${pid} [${ip}]:${port} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	app.Use(cache.New(cache.Config{
		Expiration:   10 * time.Minute,
		CacheControl: true,
	}))

	app.Get("/pokemons", handlers.ListPokemons)
	app.Get("pokemons/:id", handlers.ShowPokemon)

	app.Listen(":8080")
}
