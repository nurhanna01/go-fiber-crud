package handlers

import (
	"go-fiber-crud/models"
	"go-fiber-crud/repository"

	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	movies, err := repository.GetMovies()
	if err!=nil{
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}
	return c.JSON(movies)
}

func CreateMovie(c *fiber.Ctx)error{
	var movie models.Movies
	if err:= c.BodyParser(&movie); err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	newMovie, err := repository.CreateMovie(movie)
	if err !=nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(200).JSON(newMovie)	
}