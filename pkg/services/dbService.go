package services

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBClient() *gorm.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Println("datbase connection failed: ", err)
	}

	return db
}
