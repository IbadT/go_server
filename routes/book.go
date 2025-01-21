package routes

import (
	"github.com/IbadT/go_server/controller"
	"github.com/labstack/echo/v4"
)

func RegisterBookRoutes(e *echo.Echo) {
	bookRoute := e.Group("/book")
	bookRoute.GET("/", controller.GetAllBooks)
	bookRoute.GET("/:id", controller.GetBook)
	bookRoute.POST("/", controller.CreateBook)
	// bookRoute.PUT("/:id", controllers.UpdateBook)
	// bookRoute.DELETE("/:id", controllers.DeleteBook)
}
