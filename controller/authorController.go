package controller

import (
	"net/http"

	"github.com/IbadT/go_server/config"
	. "github.com/IbadT/go_server/model"
	"github.com/labstack/echo/v4"
)

func GetAuthor(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var author Author
	if err := db.Preload("Books").First(&author, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	return c.JSON(http.StatusOK, author)
}

func CreateAuthor(c echo.Context) error {
	author := new(Author)
	if err := c.Bind(author); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := config.DB()
	if err := db.Create(author).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create author"})
	}

	return c.JSON(http.StatusCreated, author)
}

func UpdateAuthor(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var author Author
	if err := db.First(&author, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	if err := c.Bind(&author); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := db.Save(&author).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update author"})
	}

	return c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var author Author
	if err := db.First(&author, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	if err := db.Delete(&author).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete author"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Author deleted"})
}
