package controller

import (
	"fmt"
	"net/http"

	"github.com/IbadT/go_server/config"
	"github.com/IbadT/go_server/model"
	"github.com/IbadT/go_server/service"
	"github.com/IbadT/go_server/types"

	"github.com/IbadT/go_server/utils"
	"github.com/labstack/echo/v4"
)

// работает
func GetAllUsers(c echo.Context) error {
	var users []model.User
	db := config.DB()

	if err := db.Find(&users).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Empty users")
	}

	return c.JSON(http.StatusOK, &users)
}

// работает
func GetUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User

	db := config.DB()

	if err := db.First(&user, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	// вывести только те данные, которые нужны ?!!!?
	// return c.JSON(http.StatusOK, &user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	})
}

// работает
func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	// Привязка данных из запроса
	var inputUser model.User
	if err := c.Bind(&inputUser); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	userResult, err := service.UpdateUserService(id, inputUser)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, &userResult)
}

// работает
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User
	db := config.DB()

	if err := db.First(&user, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	if err := db.Delete(&user).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	// return utils.SuccessResponse(c, http.StatusOK, user)
	return utils.SuccessResponse(c, http.StatusOK, types.Response{
		Status:  http.StatusOK,
		Message: "Success",
	})
}

// работает
func Register(c echo.Context) error {
	var user model.User

	db := config.DB()

	if err := c.Bind(&user); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	if err := user.HashPassword(user.Password); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	if err := db.Create(&user).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	return utils.SuccessResponse(c, http.StatusOK, user)
}

// работает
func Login(c echo.Context) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	var user model.User
	db := config.DB()

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	if err := user.CheckPassword(input.Password); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		fmt.Println(token)
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
	}

	return utils.SuccessResponse(c, http.StatusOK, map[string]string{
		"access_token": token,
	})
}
