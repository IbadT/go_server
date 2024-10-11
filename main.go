package main

import (
	"log"

	"github.com/IbadT/go_server/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// var users = []User{
// 	{ID: "1", Name: "Eduard", Login: "ibadtoff@gmail.com", Password: "gts530200"},
// }

func main() {
	app := fiber.New()

	// database.Connect()

	// middleware
	app.Use(logger.New())

	routes.Setup(app)

	// // xh http://localhost:3000/api
	// app.Get("/api", func(c fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	// // xh http://localhost:3000/api/Eduard
	// app.Get("/api/:name", func(c fiber.Ctx) error {
	// 	msg := fmt.Sprintf("Hello %s 👋", c.Params("name"))
	// 	return c.SendString(msg)
	// })

	// app.Get("/api/users", getUsers)

	// middleware
	// app.Use(func(c fiber.Ctx) error {
	// 	return c.SendStatus(404)
	// })

	log.Fatal(app.Listen(":3000"))
}

// go mod init github.com/your/repo
// go get -u github.com/gofiber/fiber/v3
// https://github.com/IbadT/go_server.git
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres
