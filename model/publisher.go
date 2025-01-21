package model

import "time"

type Publisher struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}
