package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config struct mirrors the structure of the YAML file
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
}

// NewConfig reads configuration from the given filename and returns a Config object
func NewConfig(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
