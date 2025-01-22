package main

import (
	"net/http"

	"github.com/IbadT/go_server/config"
	"github.com/IbadT/go_server/routes"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	e := echo.New()

	///

	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		start := time.Now()

	// 		// Выполнение следующего обработчика
	// 		err := next(c)

	// 		// Фиксация времени окончания
	// 		elapsed := time.Since(start)

	// 		// Логирование времени выполнения
	// 		c.Logger().Infof("Request took: %s", elapsed)

	// 		return err
	// 	}
	// })

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	///

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Group("/api")
	routes.RegisterRoutes(e)

	// // programmatically set swagger info
	// docs.SwaggerInfo.Title = "Swagger Example API"
	// docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	// docs.SwaggerInfo.BasePath = "/v2"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// r := gin.New()

	// // use ginSwagger middleware to serve the API docs
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Run()

	// bookRoute := e.Group("/book")
	// bookRoute.GET("/", controller.GetAllBooks)
	// bookRoute.POST("/", controller.CreateBook)
	// bookRoute.GET("/:id", controller.GetBook)
	// // bookRoute.PUT("/:id", controller.UpdateBook)
	// // bookRoute.DELETE("/:id", controller.DeleteBook)

	// authorRoute := e.Group("/author")
	// authorRoute.GET("/", controller.GetAuthor)
	// authorRoute.POST("/", controller.CreateAuthor)
	// authorRoute.PUT("/:id", controller.UpdateAuthor)
	// authorRoute.DELETE("/:id", controller.DeleteAuthor)

	// publisherRoute := e.Group("/publisher")
	// publisherRoute.GET("/", controller.GetPublisher)
	// publisherRoute.POST("/", controller.CreatePublisher)
	// publisherRoute.PUT("/:id", controller.UpdatePublisher)
	// publisherRoute.DELETE("/:id", controller.DeletePublisher)

	// categoryRoute := e.Group("/category")
	// categoryRoute.GET("/", controller.GetCategory)
	// categoryRoute.POST("/", controller.CreateCategory)
	// categoryRoute.PUT("/:id", controller.UpdateCategory)
	// categoryRoute.DELETE("/:id", controller.DeleteCategory)

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// e.Logger.Fatal(e.Start(":8080"))

	////

	// e.GET("/example", func(c echo.Context) error {
	// 	time.Sleep(100 * time.Millisecond)
	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": "Request processed successfully",
	// 	})
	// })

	e.Start(":8080")
}
