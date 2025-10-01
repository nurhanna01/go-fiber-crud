package repository

import (
	"go-fiber-crud/database"
	"go-fiber-crud/models"
)

func GetMovies() ([]models.Movies,error) {
	var movies  []models.Movies
	result := database.GetDB().Find(&movies)
	return movies, result.Error
}

func CreateMovie(movie models.Movies)(models.Movies,error){
	result := database.GetDB().Create(&movie)
	return movie, result.Error
}