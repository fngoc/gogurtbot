package openai

import (
	"context"
	"fmt"
	"gogurtbot/internal/logger"

	gpt "github.com/sashabaranov/go-openai"
)

var prompt = ``

var client *gpt.Client

const (
	gptToken = ""
	gptURL   = ""
)

// Initialize инициализация конфигурации и клиента gpt модели
func Initialize() {
	config := gpt.DefaultConfig(gptToken)
	config.BaseURL = gptURL
	client = gpt.NewClientWithConfig(config)
	logger.Log.Info("Client gpt initialized")
}

// SendMessageWithPrompt отправка сообщения в нейронную сеть
func SendMessageWithPrompt(message string) (string, error) {
	req := gpt.ChatCompletionRequest{
		Model: "",
		Messages: []gpt.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Вот текст: %s", message),
			},
		},
		Temperature: 0.7,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
