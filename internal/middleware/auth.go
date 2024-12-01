package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtKey = []byte("secret_key") // Храни ключ безопасно, используй ENV переменные

// Claims структура для токена
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT генерирует JWT токен
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT проверяет токен JWT
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// authMiddleware проверяет JWT и извлекает имя пользователя
func authMiddleware(c *gin.Context) {
	// Извлекаем токен из заголовка
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
		c.Abort()
		return
	}

	// Проверяем токен
	claims, err := ValidateJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Сохраняем имя пользователя в контексте
	c.Set("name", claims.Username)
	c.Next()
}

// protectedHandler - защищённый маршрут
func protectedHandler(c *gin.Context) {
	// Извлекаем имя пользователя из контекста
	name := c.MustGet("name").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name})
}
