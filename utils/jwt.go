package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 生成JWT令牌
func GenerateToken(username string) (string, error) {
	// 设置令牌过期时间
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	})

	// 签名令牌
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
