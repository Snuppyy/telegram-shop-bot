package utils

import (
	"encoding/json"
	"os"
)

type TelegramConfig struct {
	BotToken            string   `json:"bot_token"`
	EnableMultilanguage bool     `json:"enable_multilanguage"`
	DefaultLanguage     string   `json:"default_language"`
	AllowedLanguages    []string `json:"allowed_languages"`
}

func LoadConfig(path string) (*TelegramConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &TelegramConfig{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
