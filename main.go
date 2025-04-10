package main

import (
	"net/http"
	"old/api"
	"old/database" // 修改为正确的路径
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	// 配置CORS

	config := cors.Config{
		AllowOrigins:     []string{"*"},                                                         // 允许所有来源访问
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                   // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                            // 暴露的响应头
		AllowCredentials: true,                                                                  // 是否允许发送Cookie
		MaxAge:           12 * time.Hour,                                                        // 预检请求的缓存时间
	}

	// 初始化路由并应用CORS配置
	r := gin.Default()
	r.Use(cors.New(config))

	database.InitDB()

	// 处理GET请求，返回hello消息
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	// 初始化所有API路由
	api.InitAPIRoutes(r)

	// 启动服务器
	r.Run(":9001")
}
