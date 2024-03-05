package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	DBConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)

func Parse(fileName string) (*DBConfig, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read config from yml file: %w", err)
	}

	cfg := DBConfig{}
	err = yaml.Unmarshal(data, &cfg)
}
