package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
  gormDB *gorm.DB
  err    error
)

func InitDB() error {
	dsn := "root:@tcp(127.0.0.1:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local"
  gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    return err
  }
	return nil
}

func GetDB() *gorm.DB {
	return gormDB
}