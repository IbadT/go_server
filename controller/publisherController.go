package controller

import (
	"net/http"

	"github.com/IbadT/go_server/config"
	. "github.com/IbadT/go_server/model"
	"github.com/labstack/echo/v4"
)

func GetPublisher(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var publisher Publisher
	if err := db.Preload("Book").First(&publisher, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Publisher not found"})
	}

	return c.JSON(http.StatusOK, publisher)
}

func CreatePublisher(c echo.Context) error {
	publisher := new(Publisher)
	if err := c.Bind(publisher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := config.DB()
	if err := db.Create(publisher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create publisher"})
	}

	return c.JSON(http.StatusCreated, publisher)
}

func UpdatePublisher(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var publisher Publisher
	if err := db.First(&publisher, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Publisher not found"})
	}

	if err := c.Bind(&publisher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := db.Save(&publisher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update publisher"})
	}

	return c.JSON(http.StatusOK, publisher)
}

func DeletePublisher(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var publisher Publisher
	if err := db.First(&publisher, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Publisher not found"})
	}

	if err := db.Delete(&publisher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete publisher"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Publisher deleted"})
}
