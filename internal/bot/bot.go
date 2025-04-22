package bot

import (
	"fmt"
	"strings"

	"github.com/fngoc/gogurtbot/internal/config"
	"github.com/fngoc/gogurtbot/internal/logger"

	"github.com/mymmrac/telego"
)

var (
	bot   *telego.Bot
	queue []string
)

// Run запуск бота
func Run() error {
	botInstant, err := telego.NewBot(config.Configuration.Telegram.Token, telego.WithDefaultLogger(false, true))
	if err != nil {
		return err
	}

	bot = botInstant

	botUser, err := botInstant.GetMe()
	if err != nil {
		return err
	}

	logger.Log.Info(fmt.Sprintf("Start bot, botUser: %v", botUser))

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		return err
	}
	defer bot.StopLongPolling()
	readingMessage(updates)
	return nil
}

// readingMessage запуск чтения сообщений из чата
func readingMessage(updates <-chan telego.Update) {
	for update := range updates {
		// Проверка, на пустой ответ от telego
		if update.Message == nil {
			logger.Log.Info("Update message is empty")
			continue
		}

		chatID := update.Message.Chat.ID

		// Проверка, что бот работает в нужном чате
		if chatID != config.Configuration.Telegram.PicnicChatID {
			logger.Log.Info(
				fmt.Sprintf(
					"Attempting to use the bot in a way that is not in the picnic: %v",
					update.Message.Chat,
				))
			continue
		}

		// Проверка на команды
		switch {
		case update.Message.Text == "/what":
			if err := timeoutMiddleware(10, update, chatID, whatCommand); err != nil {
				logger.Log.Error(err.Error())
				// Если команда не выполнилась - очередь сообщений не чистим
				continue
			}
			// Очищаем очередь сообщений после успешного выполнения команды
			queue = queue[:0]
		case update.Message.Text == "/good":
			if err := timeoutMiddleware(10, update, chatID, goodCommand); err != nil {
				logger.Log.Error(err.Error())
			}
		case strings.HasPrefix(update.Message.Text, "/short"):
			if err := timeoutMiddleware(10, update, chatID, shortCommand); err != nil {
				logger.Log.Error(err.Error())
			}
		case strings.HasPrefix(update.Message.Text, "/say"):
			if err := timeoutMiddleware(10, update, chatID, sayCommand); err != nil {
				logger.Log.Error(err.Error())
			}
		default:
			// В противном случае обновляем очередь сообщений
			logger.Log.Debug(fmt.Sprintf("Received message: %s", update.Message.Text))
			updateQueue(update.Message.Text)
			debugMessage(fmt.Sprintf("Буфер сообщений: %s", formatMessage(queue)))
		}
	}
}
