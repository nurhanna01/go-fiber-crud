package handlers

import (
	"fmt"
	"go-fiber-crud/dto"
	"go-fiber-crud/models"
	"go-fiber-crud/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetMovies(c *fiber.Ctx) error {
	movies, err := repository.GetMovies()
	if err!=nil{
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}
	return c.JSON(movies)
}

func CreateMovie(c *fiber.Ctx)error{
	var req dto.CreateMovieRequest

	 // Parse body
	if err:= c.BodyParser(&req); err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}

	// validate
	if err:= validate.Struct(req);err !=nil{
		fmt.Println("Validation error",err)
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}

	movie := models.Movies{
		Title: req.Title,
		Genre: req.Genre,
		Year: req.Year,
	}

	movie, err := repository.CreateMovie(models.Movies(movie))
	if err !=nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(200).JSON(movie)	
}

func FindMovie(c*fiber.Ctx)error{
	id,err:= c.ParamsInt("id"); 
	if err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	findMovie, err := repository.DetailMovie(id)

	if err !=nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(200).JSON(findMovie)
}

func UpdateMovie(c*fiber.Ctx)error{
	var req dto.CreateMovieRequest
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error:":err.Error()})
	}
	if err := c.BodyParser(&req);
	err != nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
		// validate
	if err:= validate.Struct(req); err !=nil{
		fmt.Println("Validation error",err)
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	
	// mapping to movie DTO
	movie := models.Movies{
		Title: req.Title,
		Genre: req.Genre,
		Year: req.Year,
	}

	updateMovie,errorUpdate := repository.UpdateMovie(id,movie)
	if errorUpdate!=nil{
		return c.Status(500).JSON(fiber.Map{"error":errorUpdate.Error()})
	}
	return c.Status(200).JSON(updateMovie)

}