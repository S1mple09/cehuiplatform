package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword 验证密码
// 如果数据库密码已经加密（通常以 "$2" 开头），则使用 bcrypt 验证；
// 如果是明文，则直接比较，并可以选择将明文密码升级为加密密码（这里仅返回 true，需要在注册或登录后进行升级操作）。
func VerifyPassword(storedPassword, suppliedPassword string) bool {
	// 判断存储的密码是否为 bcrypt 哈希（bcrypt 的哈希一般以 "$2a$", "$2b$" 或 "$2y$" 开头）
	if strings.HasPrefix(storedPassword, "$2a$") ||
		strings.HasPrefix(storedPassword, "$2b$") ||
		strings.HasPrefix(storedPassword, "$2y$") {
		// 使用 bcrypt 比较哈希密码与明文密码
		err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(suppliedPassword))
		return err == nil
	}

	// 如果存储的密码不是 bcrypt 格式，认为它是明文存储，直接比较
	return storedPassword == suppliedPassword
}

// LogLoginAttempt 记录登录尝试
func LogLoginAttempt(username, ip string, duration time.Duration, success bool) {
	// 实现你自己的日志记录逻辑
	fmt.Printf("Login attempt: username=%s, ip=%s, duration=%v, success=%v\n", username, ip, duration, success)
}

// GenerateToken 生成JWT令牌
func GenerateToken(username string) (string, error) {
	// 设置令牌过期时间
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	})

	// 签名令牌（确保将 "your-secret-key" 替换为你自己的密钥）
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AuthMiddleware 定义认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "未提供认证信息"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			c.JSON(401, gin.H{"error": "无效的认证信息格式"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("无效的签名方法")
			}
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "无效的令牌"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "无效的令牌"})
			c.Abort()
			return
		}
		username := claims["username"].(string)
		c.Set("username", username)

		c.Next()
	}
}
