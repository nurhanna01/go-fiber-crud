package main

import (
	"go-fiber-crud/database"
	"go-fiber-crud/handlers"
	"go-fiber-crud/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// connect and migrate
	database.InitDB()
	database.GetDB().AutoMigrate(&models.Movies{})

	// routes
	app.Get("/movies",handlers.GetMovies)
	app.Post("/movies",handlers.CreateMovie)
	app.Get("/movies/:id",handlers.FindMovie)
	app.Put("/movies/:id",handlers.UpdateMovie)

	app.Listen(":9000")
	}