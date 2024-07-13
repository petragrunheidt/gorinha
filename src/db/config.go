package db

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
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
	env := currentEnv()
	
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

func currentEnv() string{
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Errorf("error determining current file path")
	}

	configPath := filepath.Join(filepath.Dir(currentFile), "../../.env")
	err := godotenv.Load(configPath)

	if err != nil {
		fmt.Errorf("error loading .env file: %w", err)
	}

	return os.Getenv("ENV")
}
