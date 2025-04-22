package main

import (
	"fmt"

	"github.com/fngoc/gogurtbot/internal/ai"
	"github.com/fngoc/gogurtbot/internal/bot"
	"github.com/fngoc/gogurtbot/internal/config"
	"github.com/fngoc/gogurtbot/internal/logger"
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

	ai.Initialize()

	if err := bot.Run(); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error run bot: %v", err))
	}
}
