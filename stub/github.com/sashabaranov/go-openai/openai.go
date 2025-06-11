package openai

import "context"

// Config and client stubs

type Config struct{ BaseURL string }

func DefaultConfig(token string) Config { return Config{} }

type Client struct {
	Response string
	Err      error
}

func NewClientWithConfig(cfg Config) *Client { return &Client{} }

type ChatCompletionMessage struct {
	Role    string
	Content string
}

type ChatCompletionRequest struct {
	Model       string
	Messages    []ChatCompletionMessage
	Temperature float32
}

type ChatCompletionChoice struct {
	Message ChatCompletionMessage
}

type ChatCompletionResponse struct {
	Choices []ChatCompletionChoice
}

func (c *Client) CreateChatCompletion(ctx context.Context, req ChatCompletionRequest) (ChatCompletionResponse, error) {
	if c.Err != nil {
		return ChatCompletionResponse{}, c.Err
	}
	return ChatCompletionResponse{Choices: []ChatCompletionChoice{{Message: ChatCompletionMessage{Content: c.Response}}}}, nil
}
