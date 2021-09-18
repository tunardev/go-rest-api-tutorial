package controller

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	Database *gorm.DB
}

func (c Controller) GetUsers(ctx *fiber.Ctx) error {
	var users []models.User
	c.Database.Find(&users)
	return ctx.Status(200).JSON(users)
}

func (c Controller) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User
	c.Database.Find(&user, id)
	return ctx.Status(200).JSON(user)
}

func (c Controller) NewUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
            "message": err.Error(),
		})
	}
	c.Database.Create(&user)
	return ctx.Status(200).JSON(user)
}

func (c Controller) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User
	c.Database.First(&user, id)
	if user.Username == "" {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No user found with given id",
		})
	}
	c.Database.Delete(&user)
	return ctx.SendString("Success")
}

func (c Controller) UpdateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
            "message": err.Error(),
		})
	}
	c.Database.Model(&user).Updates(user)
	return ctx.SendString("Success")
}