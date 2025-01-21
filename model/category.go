package model

import "time"

type Category struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Books     []Book    `json:"books" gorm:"many2many:book_categories;"` // Связь "многие ко многим"
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}
