package api

import (
	"net/http"

	"old/database"
	"old/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRoute(r *gin.RouterGroup) {
	r.POST("/register", func(c *gin.Context) {
		var data models.User

		// 绑定请求数据
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "无效的请求数据: " + err.Error(),
			})
			return
		}

		// 检查必填字段
		if data.Username == "" || data.Password == "" || data.Phone == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "用户名、密码和手机号不能为空",
			})
			return
		}

		// 检查用户名是否已存在
		var count int
		err := database.DB.Get(&count, "SELECT COUNT(*) FROM user WHERE username = ?", data.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "数据库查询失败: " + err.Error(),
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "用户名已存在",
			})
			return
		}

		// 检查手机号是否已存在
		err = database.DB.Get(&count, "SELECT COUNT(*) FROM user WHERE phone = ?", data.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "数据库查询失败: " + err.Error(),
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "手机号已注册",
			})
			return
		}

		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "密码加密失败: " + err.Error(),
			})
			return
		}

		// 插入数据到数据库
		query := "INSERT INTO user (username, phone, password) VALUES (?, ?, ?)"
		result, err := database.DB.Exec(query, data.Username, data.Phone, string(hashedPassword))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "注册失败: " + err.Error(),
				"details": "请检查数据库表结构",
			})
			return
		}

		// 获取插入的ID
		id, _ := result.LastInsertId()

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
			"user": gin.H{
				"id":       id,
				"username": data.Username,
				"phone":    data.Phone,
			},
		})
	})
}
