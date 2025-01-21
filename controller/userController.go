package controller

import (
	"net/http"

	"github.com/IbadT/go_server/config"
	. "github.com/IbadT/go_server/model"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "data"})
}

func CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "data"})
}

func UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "data"})
}

func DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "data"})
}

// работает
func Register(c echo.Context) error {
	var user User

	db := config.DB()

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := user.HashPassword(user.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not hash password"})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	})
}

// работает
func Login(c echo.Context) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	var user User
	db := config.DB()

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User not found",
		})
	}

	if err := user.CheckPassword(input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid password",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		// "password": user.Password,
	})
}
