package handlers

import (
	"github.com/IbadT/go_server/models"
	"github.com/gofiber/fiber/v3"
)

var users = []models.User{
	{ID: "1", Name: "Eduard", Login: "ibadtoff@gmail.com", Password: "gts530200"},
}

func Test(c fiber.Ctx) error {
	return c.SendString("Hello")
}

func GetUsers(c fiber.Ctx) error {
	// var users []models.User
	// database.DB.Find(users)
	return c.JSON(users)
}

func GetUserById(c fiber.Ctx) error {
	id := c.Params("id")
	for _, user := range users {
		if user.ID == id {
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Пользователь не найден")
}

func RegisterUser(c fiber.Ctx) error {
	return c.Send(c.Body())
}
