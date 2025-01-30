package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type BotConfig struct {
	Token               string   `yaml:"token"`
	EnableMultilanguage bool     `yaml:"enable_multilanguage"`
	DefaultLanguage     string   `yaml:"default_language"`
	AllowedLanguages    []string `yaml:"allowed_languages"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	Bot      BotConfig      `yaml:"bot"`
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (db DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.DBName, db.SSLMode,
	)
}
