package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserName string    `json:"user_name" gorm:"unique;not null"`
	Bio      string    `json:"bio"`
	Rating   float32   `json:"rating"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password"`
	RoleID   uuid.UUID `json:"role_id"`                           // Внешний ключ для связи с Role
	Role     Role      `json:"role" gorm:"foreignKey:RoleID"`     // many to one
	Avatar   Avatar    `json:"avatar" gorm:"foreignKey:UserID"`   // one to one
	Posts    []Post    `json:"posts" gorm:"foreignKey:UserID"`    // one to many
	Likes    []Like    `json:"likes" gorm:"foreignKey:UserID"`    // one to many
	Comments []Comment `json:"comments" gorm:"foreignKey:UserID"` // one to many
	// savedPosts

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
