module github.com/fngoc/gogurtbot

go 1.22.3

require (
	github.com/fsnotify/fsnotify v0.0.0-00010101000000-000000000000
	github.com/mymmrac/telego v0.0.0
	github.com/mymmrac/telego/telegoutil v0.0.0
	github.com/sashabaranov/go-openai v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v0.0.0-00010101000000-000000000000
	go.uber.org/zap v0.0.0
)

replace github.com/spf13/viper => ./stub/github.com/spf13/viper

replace github.com/fsnotify/fsnotify => ./stub/github.com/fsnotify/fsnotify

replace go.uber.org/zap => ./stub/go.uber.org/zap

replace github.com/mymmrac/telego => ./stub/github.com/mymmrac/telego

replace github.com/mymmrac/telego/telegoutil => ./stub/github.com/mymmrac/telego/telegoutil

replace github.com/sashabaranov/go-openai => ./stub/github.com/sashabaranov/go-openai

replace github.com/openai/openai-go => ./stub/github.com/openai/openai-go
