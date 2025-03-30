package api

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"old/database"
	"old/models"

	"github.com/gin-gonic/gin"
)

func getProfile(c *gin.Context) {
	username := c.MustGet("username").(string)
	
	var user models.User
	err := database.DB.Get(&user, "SELECT username, phone, avatar FROM user WHERE username = ?", username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"phone":    user.Phone,
		"avatar":   user.Avatar,
	})
}

func updateProfile(c *gin.Context) {
	username := c.MustGet("username").(string)
	
	var req struct {
		Username string `json:"username"`
		Phone    string `json:"phone"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	_, err := database.DB.Exec(
		"UPDATE user SET username = ?, phone = ? WHERE username = ?",
		req.Username, req.Phone, username,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"user": gin.H{
			"username": req.Username,
			"phone":    req.Phone,
		},
	})
}

func uploadAvatar(c *gin.Context) {
	username := c.MustGet("username").(string)
	
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传头像文件"})
		return
	}

	// 确保上传目录存在
	uploadDir := "./uploads/avatars"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	newFilename := username + "_" + time.Now().Format("20060102150405") + ext
	filePath := filepath.Join(uploadDir, newFilename)

	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存头像失败"})
		return
	}

	// 更新数据库中的头像路径
	avatarURL := "/avatars/" + newFilename
	_, err = database.DB.Exec(
		"UPDATE user SET avatar = ? WHERE username = ?",
		avatarURL, username,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新头像失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "头像上传成功",
		"avatar":  avatarURL,
	})
}