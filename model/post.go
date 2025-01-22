package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID      uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title   string    `json:"title"`
	Text    string    `json:"text"`
	Content string    `json:"content" gorm:"not null"`
	UserID  uuid.UUID `json:"user_id"` // Внешний ключ для связи с User

	// User
	// likes
	// comments
	// savedPost

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"udpated_at" gorm:"autoUpdateTime"`
}
