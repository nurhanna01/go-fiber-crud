package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
  gormDB *gorm.DB
)

func InitDB() error {
	err := godotenv.Load()
	if err !=nil {
		fmt.Println("error load ENV file")
	}

	databaseHost  := os.Getenv("DB_HOST")
	databaseUser := os.Getenv("DB_USER")
	databasePort := os.Getenv("DB_PORT")
	databasePassword := os.Getenv("DB_PASS")
	databaseName  := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",	databaseUser,databasePassword,databaseHost, databasePort,databaseName)
  gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    return err
  }
	return nil
}

func GetDB() *gorm.DB {
	return gormDB
}