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

// goodCommand запуск команды /good
func goodCommand(update telego.Update, chatID int64) error {
	logger.Log.Info(fmt.Sprintf("Comand /good start works: %v", update.Message.Chat))

	message := formatMessage(queue)
	gptAnswer, err := openai.SendRequestWithResend(message, config.Configuration.Openai.GoodPrompt)
	if err != nil {
		return fmt.Errorf("send request error: %w", err)
	}

	if err := sendToChatMessage(chatID, gptAnswer+"\n#good"); err != nil {
		logger.Log.Error(err.Error())
		if err := sendToChatMessage(chatID, fmt.Sprintf("Api нейронки выдала ошибку :(")); err != nil {
			return fmt.Errorf("error sending in chat about api error: %w", err)
		}
	}

	// Для дебага
	debugMessage(fmt.Sprintf("Был вызван /good с сообщением: %s", message))
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

// shortCommand запуск команды /short
func shortCommand(update telego.Update, chatID int64) error {
	var message string
	var gptAnswer string
	var err error

	logger.Log.Info(fmt.Sprintf("Comand /short start works: %v", update.Message.Chat))

	if update.Message != nil && update.Message.ReplyToMessage != nil {
		replyToMessage := update.Message.ReplyToMessage.Text
		gptAnswer, err = openai.SendRequestWithResend("["+replyToMessage+"]", config.Configuration.Openai.ShortPrompt)
		if err != nil {
			return fmt.Errorf("send reply message request error: %w", err)
		}
	} else {
		message = strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/short"))
		gptAnswer, err = openai.SendRequestWithResend("["+message+"]", config.Configuration.Openai.ShortPrompt)
		if err != nil {
			return fmt.Errorf("send request error: %w", err)
		}
	}

	if err = sendToChatMessage(chatID, gptAnswer+"\n#short"); err != nil {
		logger.Log.Error(err.Error())
		if err = sendToChatMessage(chatID, fmt.Sprintf("Api нейронки выдала ошибку :(")); err != nil {
			return fmt.Errorf("error sending in chat about api error: %w", err)
		}
	}

	// Для дебага
	debugMessage(fmt.Sprintf("Был вызван /short с сообщением: %s", message))
	return nil
}
