package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uuid.UUID `json:"user_id"` // Внешний ключ для связи с User
	PostID    uuid.UUID `json:"post_id"` // Внешний ключ для связи с Post
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
