package routes

import (
	"github.com/IbadT/go_server/controller"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	userRoute := e.Group("/user")
	userRoute.GET("/", controller.GetAllUsers)
	userRoute.GET("/:id", controller.GetUser)
	userRoute.POST("/login", controller.Login)
	userRoute.POST("/register", controller.Register)
	userRoute.PATCH("/:id", controller.UpdateUser)
	userRoute.DELETE("/:id", controller.DeleteUser)
}
