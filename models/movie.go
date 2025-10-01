package models

type Movies struct {
	Id    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
}