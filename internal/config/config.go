package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Config struct {
		DB DB `yaml:"db"`
	}
)

func Parse(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read config from yml file: %w", err)
	}

	cfg := Config{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config data from yml: %w", err)
	}

	return &cfg, nil
}
