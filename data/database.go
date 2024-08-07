package data

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(filepath.Join(dir, "./.env"))
	if err != nil {
		panic(err)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PW := os.Getenv("DB_PW")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", DB_USER, DB_PW, DB_HOST, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func DiscconectDb(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbSql.Close()
}
