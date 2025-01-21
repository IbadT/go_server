package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Секретный ключ для подписи токена
var jwtSecret = []byte("your-secret-key")

// GenerateToken создает JWT-токен для пользователя
func GenerateToken(userID int) (string, error) {
	// Устанавливаем claims (данные токена)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Токен действителен 24 часа
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	return token.SignedString(jwtSecret)
}
