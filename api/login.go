package api

import (
	"net/http"
	"time"

	"old/database"
	"old/middleware"
	"old/models"

	"github.com/gin-gonic/gin"
)

func LoginRoute(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数: " + err.Error(),
		})
		return
	}

	startTime := time.Now()

	// 查询用户
	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
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

	// 记录登录尝试
	middleware.LogLoginAttempt(user.Username, c.ClientIP(), time.Since(startTime), true)

	// 返回登录成功信息（不包含 JWT 令牌）
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user": gin.H{
			"username": user.Username,
			"phone":    user.Phone,
		},
	})
}
