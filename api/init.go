package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"old/middleware"
)

func InitAPIRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	
	// 公共API路由(不需要认证)
	publicGroup := apiGroup.Group("")
	{
		RegisterRoute(publicGroup)
		LoginRoute(publicGroup)
	}
	
	// 需要认证的API路由
	authGroup := apiGroup.Group("")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(http.StatusOK, gin.H{"message": "欢迎 " + username})
		})
	}
}
