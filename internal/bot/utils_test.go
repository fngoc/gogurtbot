package bot

import (
	"testing"
	"time"

	"github.com/fngoc/gogurtbot/internal/config"
	"github.com/mymmrac/telego"
)

func setup() {
	config.Configuration.Telegram.MaxQueueSize = 3
	config.Configuration.Telegram.MaxMessageSize = 5
	config.Configuration.Telegram.DebugChatID = 1
	bot = &telego.Bot{}
	queue = nil
}

func TestUpdateQueue(t *testing.T) {
	setup()

	updateQueue("hello")
	if len(queue) != 1 || queue[0] != "hello" {
		t.Fatalf("queue not updated: %v", queue)
	}

	updateQueue("toolongtext")
	if len(queue) != 2 || queue[1] != "toolo" { // truncated to 5 chars
		t.Fatalf("queue truncation failed: %v", queue)
	}

	updateQueue("2")
	updateQueue("3")
	expected := []string{"toolo", "2", "3"}
	for i, v := range expected {
		if queue[i] != v {
			t.Fatalf("expected %v got %v", expected, queue)
		}
	}
}

func TestFormatMessage(t *testing.T) {
	msgs := []string{"a", "b"}
	got := formatMessage(msgs)
	want := "[(a), (b), ]"
	if got != want {
		t.Fatalf("expected %q got %q", want, got)
	}
}

func TestTimeoutMiddleware(t *testing.T) {
	setup()
	chatID := int64(123)
	executed := false
	cmd := func(u telego.Update, c int64) error { executed = true; return nil }

	lastRequestTime = time.Now().Add(-time.Hour)
	if err := timeoutMiddleware(10, telego.Update{Message: &telego.Message{}}, chatID, cmd); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !executed {
		t.Fatal("command not executed")
	}

	executed = false
	if err := timeoutMiddleware(10, telego.Update{Message: &telego.Message{}}, chatID, cmd); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if executed {
		t.Fatal("command executed despite timeout")
	}
	if len(bot.Sent) == 0 {
		t.Fatal("no warning message sent")
	}
}
