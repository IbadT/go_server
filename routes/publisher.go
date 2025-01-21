package routes

import (
	"github.com/IbadT/go_server/controller"
	"github.com/labstack/echo/v4"
)

func RegisterPublisherRoutes(e *echo.Echo) {
	publisherRoute := e.Group("/publisher")
	publisherRoute.GET("/", controller.GetPublisher)
	publisherRoute.POST("/", controller.CreatePublisher)
	publisherRoute.PUT("/:id", controller.UpdatePublisher)
	publisherRoute.DELETE("/:id", controller.DeletePublisher)
}
