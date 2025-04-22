package bot

import (
	"fmt"
	"strings"

	"github.com/fngoc/gogurtbot/internal/config"
	"github.com/fngoc/gogurtbot/internal/logger"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// sendToChatMessage отправка сообщения в чат
func sendToChatMessage(chatID int64, text string) error {
	_, err := bot.SendMessage(
		tu.Message(
			telego.ChatID{
				ID: chatID,
			},
			text,
		),
	)
	return err
}

// debugMessage функция для отправки дебаг сообщений в чат
func debugMessage(message string) {
	if err := sendToChatMessage(
		config.Configuration.Telegram.DebugChatID,
		message,
	); err != nil {
		logger.Log.Error(err.Error())
	}
}

// updateQueue обновления очереди сообщений исходя из заданной максимальной длинны
func updateQueue(text string) {
	if len(text) == 0 {
		return
	}
	if len(text) > config.Configuration.Telegram.MaxMessageSize {
		text = text[:config.Configuration.Telegram.MaxMessageSize]
	}

	if len(queue) >= config.Configuration.Telegram.MaxQueueSize {
		queue = append(queue[1:], text)
	} else {
		queue = append(queue, text)
	}
}

// formatMessage форматирование очереди сообщений для нейронной сети
func formatMessage(q []string) string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, val := range q {
		sb.WriteString(fmt.Sprintf("(%s), ", val))
	}
	sb.WriteString("]")
	return sb.String()
}
