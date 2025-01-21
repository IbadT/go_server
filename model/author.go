package model

import "time"

type Author struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Books     []Book    `json:"books" gorm:"foreignKey:AuthorID"` // Связь "один ко многим"
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
