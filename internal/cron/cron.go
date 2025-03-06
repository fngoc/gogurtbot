package cron

import "github.com/robfig/cron/v3"

func BirthdaysScheduler() error {
	c := cron.New()
	_, err := c.AddFunc("0 9 * * *", func() {}) //todo добавить проверку, что у кого-то день рождения и отправку сообщения с поздравлением
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
