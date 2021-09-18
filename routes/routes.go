package routes

import (
	"go-rest-api/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, controllers controller.Controller) {

	app.Get("/api", controllers.GetUsers)
	app.Get("/api/:id", controllers.GetUser)
	app.Post("/api", controllers.NewUser)
	app.Delete("/api/:id", controllers.DeleteUser)
	app.Put("/api/:id", controllers.UpdateUser)

}