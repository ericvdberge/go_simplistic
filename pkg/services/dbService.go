package services

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBClient() *gorm.DB {
	dsn := "host=database user=postgres password=postgres dbname=app port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("datbase connection failed: ", err)
	}

	return db

}
