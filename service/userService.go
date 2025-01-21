package service

import (
	"github.com/IbadT/go_server/config"
	"github.com/IbadT/go_server/model"
	"github.com/IbadT/go_server/utils"
)

func UpdateUserService(id string, inputUser model.User) (*model.User, error) {
	var user model.User
	db := config.DB()

	if err := db.First(&user, id).Error; err != nil {
		return nil, utils.NewError("user not found")
	}

	var existingUser model.User
	if err := db.Where("email = ?", inputUser.Email).First(&existingUser).Error; err != nil {
		if existingUser.ID != user.ID {
			return nil, utils.NewError("email is already in use")
		}
	}

	if err := inputUser.HashPassword(inputUser.Password); err != nil {
		return nil, utils.NewError("failed to hash password")
	}

	if err := db.Model(&user).Updates(inputUser).Error; err != nil {
		return nil, utils.NewError("failed to update user")
	}

	return &user, nil

}
