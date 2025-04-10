package api

import (
	"github.com/gin-gonic/gin"
	"old/middleware"
)

func InitAPIRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	// 公共API路由(不需要认证)
	publicGroup := apiGroup.Group("")
	{
		publicGroup.POST("/register", RegisterRoute)
		publicGroup.POST("/login", LoginRoute)
	}

	// 需要认证的API路由
	authGroup := apiGroup.Group("")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(200, gin.H{"message": "欢迎 " + username})
		})
		authGroup.POST("/scan", func(c *gin.Context) { //扫描时候必需提供目标
			url := c.MustGet("url").(string)
			c.JSON(200, gin.H{"url": url})

		})
	}
}
