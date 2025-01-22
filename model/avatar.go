package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Avatar struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Avatar   []byte    `json:"avatar" gorm:"type:bytea;not null"`         // Бинарные данные (аналог bytea в PostgreSQL)
	Mimetype string    `json:"mimetype" gorm:"type:varchar;default:null"` // Тип MIME (например, image/jpeg)
	Filename string    `json:"filename" gorm:"type:varchar;default:null"` // Имя файла

	UserID uuid.UUID `json:"user_id" gorm:"unique;not null"` // Внешний ключ для связи с User
	// User   User      `json:"user" gorm:"foreignKey:UserID"`  // Связь One-to-One

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Автоматически создается
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Автоматически обновляется
}
