package db

import (
	db "test/pkg/services"
)

type Todo struct {
	Id      int    `gorm:"primaryKey"`
	Message string `json:"message"`
}

func Migrate() {
	db := db.GetDBClient()
	db.AutoMigrate(&Todo{})
}
