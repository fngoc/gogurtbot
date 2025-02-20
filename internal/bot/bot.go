package bot

import (
	"fmt"
	"gogurtbot/internal/config"
	"gogurtbot/internal/logger"
	"gogurtbot/internal/openai"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var (
	bot             *telego.Bot
	queue           []string
	lastRequestTime time.Time
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
			logger.Log.Info(fmt.Sprintf("Update message is empty"))
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

		// Если отправлена команда
		if update.Message.Text == "/what" {
			currentTime := time.Now()

			// Проверяем, прошло ли 30 секунд с последнего запроса
			if currentTime.Sub(lastRequestTime) >= 30*time.Second {
				lastRequestTime = currentTime

				logger.Log.Info(fmt.Sprintf("Comand /what start works: %v", update.Message.Chat))

				message := formatMessage(queue)
				gptAnswer, err := openai.SendRequestWithResend(message)
				if err != nil {
					logger.Log.Error(fmt.Sprintf("Error sending chat: %v", err))
				}

				if err := sendToChatMessage(bot, chatID, gptAnswer); err != nil {
					logger.Log.Error(err.Error())
					if err := sendToChatMessage(bot, chatID,
						fmt.Sprintf("Api нейронки выдала ошибку :("),
					); err != nil {
						logger.Log.Error(fmt.Sprintf("Error sending in debug chat: %v", err))
					}
					// Если команда не выполнилась - очередь сообщений не чистим
					continue
				}

				// Очищаем очередь сообщений после успешного выполнения команды
				queue = queue[:0]
			} else {
				// Если прошло меньше 30 секунд
				logger.Log.Info("Too many request")
				if err := sendToChatMessage(bot, chatID,
					fmt.Sprintf("Слишком много запросов, подожди 30 секунд"),
				); err != nil {
					logger.Log.Error(err.Error())
				}
			}
		} else {
			// В противном случае обновляем очередь сообщений
			updateQueue(update.Message.Text)
			logger.Log.Info("Received message: " + update.Message.Text)
			// Для дебага
			if err := sendToChatMessage(
				bot, config.Configuration.Telegram.DebugChatID,
				fmt.Sprintf("Буфер сообщений: %s", formatMessage(queue)),
			); err != nil {
				logger.Log.Error(err.Error())
			}
		}
	}
}

// sendToChatMessage отправка сообщения в чат
func sendToChatMessage(bot *telego.Bot, chatID int64, text string) error {
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
