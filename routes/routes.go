package routes

import (
	"github.com/IbadT/go_server/handlers"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {

	// xh post http://localhost:3000/api/
	app.Get("/api", handlers.Test)

	// xh post http://localhost:3000/api/users
	app.Get("/api/users", handlers.GetUsers)

	// xh post http://localhost:3000/api/user/1
	app.Get("/api/user/:id", handlers.GetUserById)

	// xh post http://localhost:3000/api/register id=2 name=test login=it password=123
	app.Post("/api/register", handlers.RegisterUser)
}
