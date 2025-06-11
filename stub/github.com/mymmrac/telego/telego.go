package telego

// Minimal stubs for testing

type Bot struct {
	Sent []string
}

type ChatID struct{ ID int64 }

type Chat struct{ ID int64 }

type Message struct {
	Text           string
	Chat           Chat
	ReplyToMessage *Message
}

type Update struct {
	Message *Message
}

type User struct{}

func NewBot(token string, opts ...interface{}) (*Bot, error) { return &Bot{}, nil }
func WithDefaultLogger(b1, b2 bool) interface{}              { return nil }

func (b *Bot) GetMe() (User, error) { return User{}, nil }
func (b *Bot) UpdatesViaLongPolling(v interface{}) (<-chan Update, error) {
	return make(chan Update), nil
}
func (b *Bot) StopLongPolling() {}

type MessageParams struct {
	ChatID ChatID
	Text   string
}

func (b *Bot) SendMessage(p interface{}) (interface{}, error) {
	switch v := p.(type) {
	case MessageParams:
		b.Sent = append(b.Sent, v.Text)
	case *MessageParams:
		b.Sent = append(b.Sent, v.Text)
	}
	return nil, nil
}
