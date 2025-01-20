package models

type User struct {
	// gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey;autoincrement"`
	Name     string `json:"name"`
	Login    string `json:"login" gorm:"unique"`
	Password string `json:"password"`
}
