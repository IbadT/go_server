package routes

import (
	"github.com/IbadT/go_server/controller"
	"github.com/labstack/echo/v4"
)

func RegisterCategoryRoutes(e *echo.Echo) {
	categoryRoute := e.Group("/category")
	categoryRoute.GET("/", controller.GetCategory)
	categoryRoute.POST("/", controller.CreateCategory)
	categoryRoute.PUT("/:id", controller.UpdateCategory)
	categoryRoute.DELETE("/:id", controller.DeleteCategory)
}
