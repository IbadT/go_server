package routes

import (
	"github.com/IbadT/go_server/controller"
	"github.com/labstack/echo/v4"
)

func RegisterAuthorRoutes(e *echo.Echo) {
	authorRoute := e.Group("/author")
	authorRoute.GET("/", controller.GetAuthor)
	authorRoute.POST("/", controller.CreateAuthor)
	authorRoute.PUT("/:id", controller.UpdateAuthor)
	authorRoute.DELETE("/:id", controller.DeleteAuthor)
}
