package router

import (
	"GoCodeEval/internal/config"
	"GoCodeEval/internal/handler"
	"GoCodeEval/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router with the specified routes.
func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// 使用中间件
	router.Use(middleware.ExampleMiddleware())

	// 设置路由组
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", handler.RegisterHandler)
		}
	}

	return router
}
