package ai

import (
	"testing"

	"github.com/fngoc/gogurtbot/internal/config"
	openai "github.com/sashabaranov/go-openai"
)

func TestSendRequestWithResend(t *testing.T) {
	config.Configuration.Ai.CountRepeatedRequests = 2
	client = &openai.Client{Response: "answer"}

	resp, err := SendRequestWithResend("msg", "prompt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp != "answer" {
		t.Fatalf("expected answer got %q", resp)
	}
}

func TestSendMessageWithPrompt(t *testing.T) {
	client = &openai.Client{Response: "world"}
	config.Configuration.Ai.Model = "model"
	ans, err := SendMessageWithPrompt("hello", "prompt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ans != "world" {
		t.Fatalf("expected world got %q", ans)
	}
}
