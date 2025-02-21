package bot

import (
	"fmt"
	"gogurtbot/internal/config"
	"gogurtbot/internal/logger"
	"gogurtbot/internal/openai"
	"strings"

	"github.com/mymmrac/telego"
)

// whatCommand запуск команды /what
func whatCommand(update telego.Update, chatID int64) error {
	logger.Log.Info(fmt.Sprintf("Comand /what start works: %v", update.Message.Chat))

	message := formatMessage(queue)
	gptAnswer, err := openai.SendRequestWithResend(message, config.Configuration.Openai.WhatPrompt)
	if err != nil {
		return fmt.Errorf("send request error: %w", err)
	}

	if err := sendToChatMessage(chatID, gptAnswer+"\n#what"); err != nil {
		logger.Log.Error(err.Error())
		if err := sendToChatMessage(chatID, fmt.Sprintf("Api нейронки выдала ошибку :(")); err != nil {
			return fmt.Errorf("error sending in chat about api error: %w", err)
		}
	}

	// Для дебага
	debugMessage(fmt.Sprintf("Был вызван /what с сообщением: %s", message))
	return nil
}

// sayCommand запуск команды /say
func sayCommand(update telego.Update, chatID int64) error {
	logger.Log.Info(fmt.Sprintf("Comand /say start works: %v", update.Message.Chat))

	message := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/say"))
	gptAnswer, err := openai.SendRequestWithResend("["+message+"]", config.Configuration.Openai.SayPrompt)
	if err != nil {
		return fmt.Errorf("send request error: %w", err)
	}

	if err := sendToChatMessage(chatID, gptAnswer+"\n#say"); err != nil {
		logger.Log.Error(err.Error())
		if err := sendToChatMessage(chatID, fmt.Sprintf("Api нейронки выдала ошибку :(")); err != nil {
			return fmt.Errorf("error sending in chat about api error: %w", err)
		}
	}

	// Для дебага
	debugMessage(fmt.Sprintf("Был вызван /say с сообщением: %s", message))
	return nil
}
