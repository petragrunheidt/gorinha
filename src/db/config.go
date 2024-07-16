package db

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func LoadConfig(filename string) (*DBConfig, error) {
	env := gin.Mode()
	
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	var config map[string]DBConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	dbConfig, ok := config[env]
	if !ok {
		return nil, fmt.Errorf("configuration for environment %s not found", env)
	}

	return &dbConfig, nil
}
