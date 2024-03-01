package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func LoadDatabaseConfig() (DatabaseConfig, error) {
	var config DatabaseConfig

	// 读取 YAML 文件内容
	data, err := ioutil.ReadFile("internal/config/database.yaml")
	if err != nil {
		return config, err
	}

	// 解析 YAML 数据到结构体
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
