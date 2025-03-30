package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("zhangyuxiangdefofashanzhaizhilu") // 请替换为安全的密钥

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyPassword(hashedPassword, inputPassword string) bool {
	// 这里应该实现密码验证逻辑，比如bcrypt比较
	// 暂时简单比较，实际项目中应该使用安全的方式
	return hashedPassword == inputPassword
}

func LogLoginAttempt(username, ip string, duration time.Duration, success bool) {
	// 这里应该实现登录日志记录逻辑
	// 可以记录到数据库或日志文件中
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供授权令牌"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的授权令牌"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}