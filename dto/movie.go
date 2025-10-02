package dto

type CreateMovieRequest struct {
	Title string `json:"title" validate:"required"`
	Genre string `json:"genre"`
	Year  int    `json:"year" validate:"required"`
}