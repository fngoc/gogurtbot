package config

import (
	"fmt"
	"gogurtbot/internal/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config структура конфигурации
type Config struct {
	Telegram struct {
		Token          string `mapstructure:"token"`
		PicnicChatID   int64  `mapstructure:"picnicChatID"`
		DebugChatID    int64  `mapstructure:"debugChatID"`
		MaxMessageSize int    `mapstructure:"maxMessageSize"`
		MaxQueueSize   int    `mapstructure:"maxQueueSize"`
	} `mapstructure:"telegram"`

	Ai struct {
		CountRepeatedRequests int    `mapstructure:"countRepeatedRequests"`
		Token                 string `mapstructure:"token"`
		Model                 string `mapstructure:"model"`
		GptURL                string `mapstructure:"gptURL"`
		WhatPrompt            string `mapstructure:"whatPrompt"`
		GoodPrompt            string `mapstructure:"goodPrompt"`
		SayPrompt             string `mapstructure:"sayPrompt"`
		ShortPrompt           string `mapstructure:"shortPrompt"`
	} `mapstructure:"ai"`
}

// Configuration инстант конфигурации
var Configuration Config

// LoadConfig парсит конфигурационный yaml файл
func LoadConfig() error {
	viper.SetConfigName("config") // имя файла конфигурации без расширения
	viper.SetConfigType("yaml")   // тип файла конфигурации
	viper.AddConfigPath(".")      // путь к файлу конфигурации (текущая директория)

	// Чтение конфигурационного файла
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	// Инициализация структуры конфигурации
	if err := viper.Unmarshal(&Configuration); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Настройка отслеживания изменений в конфигурационном файле
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Log.Info(fmt.Sprintf("Config file changed: %s", e.Name))
		if err := viper.Unmarshal(&Configuration); err != nil {
			logger.Log.Info(fmt.Sprintf("Failed to unmarshal config: %v", err))
		} else {
			logger.Log.Info(fmt.Sprintf("Reloaded config"))
		}
	})

	logger.Log.Info("Loaded config")
	return nil
}
