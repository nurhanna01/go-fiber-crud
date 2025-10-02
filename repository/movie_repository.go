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

func DetailMovie(id int)(models.Movies,error){
	var movies models.Movies
	result := database.GetDB().First(&movies,"id = ?",id)
	return movies, result.Error
}

func UpdateMovie(id int,movie models.Movies)(models.Movies,error){
	var existingMovies models.Movies
	findMovie := database.GetDB().First(&existingMovies,"id=?",id)
	if findMovie.Error!=nil{
		return movie,findMovie.Error
	}
	existingMovies.Title=movie.Title
	existingMovies.Genre=movie.Genre
	existingMovies.Year=movie.Year
	updateResult := database.GetDB().Save(&existingMovies)
	return existingMovies,updateResult.Error
}