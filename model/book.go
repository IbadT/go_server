package model

import "time"

type Book struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"` // Время создания записи
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Время последнего обновления записи
}
