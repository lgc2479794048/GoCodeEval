package main

import (
	"GoCodeEval/internal/config" // 更新为你的项目路径
	"GoCodeEval/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	cfg, err := config.NewConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Load config failed: %v", err)
	}

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// // 初始化Gin引擎
	// r := gin.Default()

	// ... 设置路由等
	// 设置路由
	r := router.SetupRouter(cfg)
	// 运行服务
	r.Run(cfg.Server.Port) // listen and serve on the port specified in the config
}
