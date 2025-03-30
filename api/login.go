package api

import (
	"net/http"
	"old/database"
	"old/middleware"
	"old/models"
	"time"

	"github.com/gin-gonic/gin"
)

// LoginRequest 定义登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// LoginResponse 定义登录响应结构体
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    struct {
		Username string `json:"username"`
		Phone    string `json:"phone"`
	} `json:"user"`
}

func LoginRoute(r *gin.RouterGroup) {
	r.POST("/login", func(c *gin.Context) {
		time.Now()

		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "无效的请求参数: " + err.Error(),
			})
			return
		}

		startTime := time.Now()

		// 查询用户
		var user models.User
		query := "SELECT id, username, phone, password FROM user WHERE username = ? LIMIT 1"
		err := database.DB.Get(&user, query, req.Username)
		if err != nil {
			middleware.LogLoginAttempt(req.Username, c.ClientIP(), time.Since(startTime), false)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "用户名或密码错误",
			})
			return
		}

		// 验证密码
		if !middleware.VerifyPassword(user.Password, req.Password) {
			middleware.LogLoginAttempt(user.Username, c.ClientIP(), time.Since(startTime), false)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "用户名或密码错误",
			})
			return
		}

		// 生成JWT令牌
		token, err := middleware.GenerateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "生成令牌失败",
			})
			return
		}

		// 构造响应
		resp := LoginResponse{
			Message: "登录成功",
			Token:   token,
			User: struct {
				Username string `json:"username"`
				Phone    string `json:"phone"`
			}{
				Username: user.Username,
				Phone:    user.Phone,
			},
		}

		// 设置Authorization头
		c.Header("Authorization", "Bearer "+token)

		// 记录登录日志
		middleware.LogLoginAttempt(user.Username, c.ClientIP(), time.Since(startTime), true)

		c.JSON(http.StatusOK, resp)
	})
}
