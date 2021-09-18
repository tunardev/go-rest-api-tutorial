package database

import (
	"database/sql"
	"fmt"
	"go-rest-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=postgres password=postgretunar2000 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection succesfully opened")

	db.AutoMigrate(&models.User{})

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	return db, sqlDb;
}