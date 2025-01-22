package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model

	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RequesterID uuid.UUID `json:"requester_id" gorm:"type:uuid;not null"`  // Внешний ключ для Requester
	Requester   User      `json:"requester" gorm:"foreignKey:RequesterID"` // Связь Many-to-One
	AddresseeID uuid.UUID `json:"addressee_id" gorm:"type:uuid;not null"`  // Внешний ключ для Addressee
	Addressee   User      `json:"addressee" gorm:"foreignKey:AddresseeID"` // Связь Many-to-One
	IsAccepted  bool      `json:"is_accepted" gorm:"default:false"`        // Статус принятия запроса
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`        // Автоматически создается
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`        // Автоматически обновляется
}
