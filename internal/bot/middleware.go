package bot

import (
	"fmt"
	"time"

	"github.com/fngoc/gogurtbot/internal/logger"

	"github.com/mymmrac/telego"
)

var lastRequestTime time.Time

// timeoutMiddleware мидлвар для таймаута запросов
func timeoutMiddleware(
	secondsTimeout int64,
	update telego.Update,
	chatID int64,
	command func(update telego.Update, chatID int64) error,
) error {
	currentTime := time.Now()

	// Проверяем, прошло ли secondsTimeout секунд с последнего запроса
	if currentTime.Sub(lastRequestTime) >= time.Duration(secondsTimeout)*time.Second {
		lastRequestTime = currentTime
		if err := command(update, chatID); err != nil {
			return err
		}
	} else {
		// Если прошло меньше secondsTimeout секунд
		logger.Log.Info("Too many request")
		if err := sendToChatMessage(
			chatID,
			fmt.Sprintf("Слишком много запросов, подожди %d секунд", lastRequestTime.Second()),
		); err != nil {
			logger.Log.Error(err.Error())
		}
	}
	return nil
}
