package model

import "time"

type Book struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	AuthorID    int        `json:"author_id"` // Внешний ключ для связи "один ко многим"
	Author      Author     `json:"author" gorm:"foreignKey:AuthorID"`
	PublisherID int        `json:"publisher_id"` // Внешний ключ для связи "один к одному"
	Publisher   Publisher  `json:"publisher" gorm:"foreignKey:PublisherID"`
	Categories  []Category `json:"categories" gorm:"many2many:book_categories;"` // Связь "многие ко многим"
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

// type Author struct {
// 	ID        int       `gorm:"primaryKey" json:"id"`
// 	Name      string    `json:"name"`
// 	Books     []Book    `json:"books" gorm:"foreignKey:AuthorID"` // Связь "один ко многим"
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
// }

//	type Publisher struct {
//		ID        int       `gorm:"primaryKey" json:"id"`
//		Name      string    `json:"name"`
//		Book      Book      `json:"book" gorm:"foreignKey:PublisherID"` // Связь "один к одному"
//		CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
//		UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
//	}

// type Publisher struct {
// 	ID        int       `gorm:"primaryKey" json:"id"`
// 	Name      string    `json:"name"`
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
// }

// type Category struct {
// 	ID        int       `gorm:"primaryKey" json:"id"`
// 	Name      string    `json:"name"`
// 	Books     []Book    `json:"books" gorm:"many2many:book_categories;"` // Связь "многие ко многим"
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
// }
