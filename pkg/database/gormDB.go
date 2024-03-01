package database

import (
	"GoCodeEval/internal/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitializeDB 创建并返回一个新的数据库连接
func InitializeDB() (*gorm.DB, error) {
	cfg, err := config.LoadDatabaseConfig()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// ... 自动迁移模型和其他初始化逻辑 ...

	return db, nil
}
