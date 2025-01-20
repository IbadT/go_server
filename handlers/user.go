package handlers

import (
	"fmt"
	"log"

	"github.com/IbadT/go_server/database"
	"github.com/IbadT/go_server/models"
	"github.com/gofiber/fiber/v3"
)

// var users = []models.User{
// 	{ID: "1", Name: "Eduard", Login: "ibadtoff@gmail.com", Password: "gts530200"},
// }

func Test(c fiber.Ctx) error {
	return c.SendString("Hello")
}

func GetUsers(c fiber.Ctx) error {
	// var users []models.User
	// database.DB.Find(users)

	return c.JSON("")
}

func GetUserById(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DB.First(user, id)
	if result.Error != nil {
		c.Status(404).SendString(result.Error.Error())
		log.Fatal("Пользователь не найден")
	}
	return c.JSON(user)
	// for _, user := range users {
	// 	if user.ID == id {
	// 		return c.JSON(user)
	// 	}
	// }
	// return c.Status(fiber.StatusNotFound).SendString("Пользователь не найден")
}

// var user = models.User{ID: "1", Name: "Eduard", Login: "ibadtoff@gmail.com", Password: "gts530200"}
func RegisterUser(c fiber.Ctx) error {
	// user := new(models.User)
	var user models.User
	// u := c.Bind().Body(&user)
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	// result := database.DB.Create(models.User{ID: user.ID, Name: user.Name, Login: user.Login, Password: user.Password})
	fmt.Println(&user)
	result := database.DB.Create(user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	// {
	// 	"id": 0,
	// 	"name": "Eduard",
	// 	"login": "ibadtoff@gmail.com",
	// 	"password": "gts530200"
	// }
	return c.Status(fiber.StatusCreated).JSON(user)
}
