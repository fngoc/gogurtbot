package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gogurtbot/internal/logger"
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

	Openai struct {
		CountRepeatedRequests int    `mapstructure:"countRepeatedRequests"`
		Token                 string `mapstructure:"token"`
		GptURL                string `mapstructure:"gptURL"`
		Prompt                string `mapstructure:"prompt"`
	}
}

// Configuration инстант конфигурации
var Configuration *Config

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
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	Configuration = &config
	logger.Log.Info("Loaded config")
	return nil
}
