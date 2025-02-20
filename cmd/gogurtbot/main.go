package main

import (
	"fmt"
	"gogurtbot/internal/bot"
	"gogurtbot/internal/config"
	"gogurtbot/internal/logger"
	"gogurtbot/internal/openai"
)

// main старт программы
func main() {
	config.ParseFlag()

	if err := logger.Initialize(config.Loglevel); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error logger initialize: %v", err))
	}

	if err := config.LoadConfig(); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error reading config file: %s", err))
	}

	openai.Initialize()

	if err := bot.Run(); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error run bot: %v", err))
	}
}
