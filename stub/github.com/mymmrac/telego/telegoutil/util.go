package telegoutil

import "github.com/mymmrac/telego"

func Message(chatID telego.ChatID, text string) telego.MessageParams {
	return telego.MessageParams{ChatID: chatID, Text: text}
}
