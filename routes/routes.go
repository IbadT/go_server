package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	RegisterBookRoutes(e)
	RegisterAuthorRoutes(e)
	RegisterPublisherRoutes(e)
	RegisterCategoryRoutes(e)
	RegisterUserRoutes(e) // Новый роут
}
