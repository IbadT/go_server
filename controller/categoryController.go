package controller

import (
	"net/http"

	"github.com/IbadT/go_server/config"
	. "github.com/IbadT/go_server/model"
	"github.com/labstack/echo/v4"
)

func GetCategory(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var category Category
	if err := db.Preload("Books").First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	category := new(Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := config.DB()
	if err := db.Create(category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create category"})
	}

	return c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var category Category
	if err := db.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := db.Save(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update category"})
	}

	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var category Category
	if err := db.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	if err := db.Delete(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete category"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted"})
}
