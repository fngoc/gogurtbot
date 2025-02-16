package main

import (
	"fmt"
	"gogurtbot/internal/bot"
	"gogurtbot/internal/logger"
	"gogurtbot/internal/openai"
)

// main старт программы
func main() {
	if err := logger.Initialize(); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error logger initialize: %v", err))
	}

	openai.Initialize()

	if err := bot.Run(); err != nil {
		logger.Log.Fatal(fmt.Sprintf("Error run bot: %v", err))
	}
}
