package api

import (
	"net/http"
	"old/database"
	"old/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRoute 用户注册接口
func RegisterRoute(c *gin.Context) {
	var data models.User

	// 绑定请求数据
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数: " + err.Error(),
		})
		return
	}

	// 检查必填字段是否为空
	if data.Username == "" || data.Password == "" || data.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名、密码和手机号不能为空",
		})
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := database.DB.Where("username = ?", data.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名已存在",
		})
		return
	}

	// 检查手机号是否已注册
	if err := database.DB.Where("phone = ?", data.Phone).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "手机号已注册",
		})
		return
	}

	// 加密密码（注册时必须使用加密存储，避免后续登录时的兼容问题）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "密码加密失败: " + err.Error(),
		})
		return
	}
	data.Password = string(hashedPassword)

	// 创建用户数据
	if err := database.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "注册失败: 数据库操作错误",
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"user": gin.H{
			"id":       data.ID,
			"username": data.Username,
			"phone":    data.Phone,
		},
	})
}
