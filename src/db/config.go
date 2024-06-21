package db

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"db"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return ParseYAML(data)
}

func ParseYAML(yamlData []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(yamlData, &config)
	if err != nil {
			return nil, fmt.Errorf("error unmarshaling config file: %w", err)
	}

	return &config, nil
}

