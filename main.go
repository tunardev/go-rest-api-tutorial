package main

import (
	"fmt"
	"go-rest-api/controller"
	"go-rest-api/database"
	"go-rest-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {
	app := fiber.New()

	db, sqlDb := database.ConnectDb()
	defer sqlDb.Close() 

	controllers := controller.Controller{ Database: db }
	routes.Setup(app, controllers)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello")
	})

	app.Listen(":8080")
}