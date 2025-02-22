package ai

import (
	"context"
	"fmt"
	"gogurtbot/internal/config"
	"gogurtbot/internal/logger"
	"math/rand"
	"time"

	gpt "github.com/sashabaranov/go-openai"
)

var client *gpt.Client

// Initialize инициализация конфигурации и клиента gpt модели
func Initialize() {
	conf := gpt.DefaultConfig(config.Configuration.Ai.Token)
	conf.BaseURL = config.Configuration.Ai.GptURL
	client = gpt.NewClientWithConfig(conf)
	logger.Log.Info("Client gpt initialized")
}

// SendRequestWithResend отправка запроса с докатом в случае ошибки со стороны api
func SendRequestWithResend(message, prompt string) (string, error) {
	for i := 0; i < config.Configuration.Ai.CountRepeatedRequests; i++ {
		answer, err := SendMessageWithPrompt(message, prompt)
		randomSecond := rand.Intn(5)
		timeSleep := time.Duration(randomSecond * int(time.Second))

		if err != nil {
			logger.Log.Warn(
				fmt.Sprintf(
					"Error sending message: %v attempt to resend the request via %ds, we'll repeat %d more times",
					err, timeSleep/time.Second, config.Configuration.Ai.CountRepeatedRequests-i,
				),
			)
		} else if answer == "" {
			logger.Log.Warn(
				fmt.Sprintf(
					"Empty answer from ML api, attempt to resend the request via %ds we'll repeat %d more times",
					timeSleep/time.Second, config.Configuration.Ai.CountRepeatedRequests-i,
				),
			)
		} else {
			return answer, nil
		}

		time.Sleep(timeSleep)
	}
	return "", nil
}

// SendMessageWithPrompt отправка сообщения в нейронную сеть
func SendMessageWithPrompt(message, prompt string) (string, error) {
	logger.Log.Debug(fmt.Sprintf("Sending message: %s", message))
	req := gpt.ChatCompletionRequest{
		Model: config.Configuration.Ai.Model,
		Messages: []gpt.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: fmt.Sprintf(" Вот текст: %s", message),
			},
		},
		Temperature: 0.7,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)

	if err != nil {
		return "", err
	}

	logger.Log.Debug(fmt.Sprintf("Full response: %+v", resp.Choices[0].Message))
	return resp.Choices[0].Message.Content, nil
}
