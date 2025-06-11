package config

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	dir := t.TempDir()
	content := `{"telegram":{"debugChatID":123,"maxMessageSize":5,"maxQueueSize":3,"token":"t","picnicChatID":1},"ai":{"countRepeatedRequests":2,"token":"a","model":"m","gptURL":"u"}}`
	if err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(content), 0644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(dir)
	if err := LoadConfig(); err != nil {
		t.Fatalf("load config: %v", err)
	}
	if Configuration.Telegram.DebugChatID != 123 {
		t.Fatalf("unexpected DebugChatID %d", Configuration.Telegram.DebugChatID)
	}
}

func TestLoadConfigMissing(t *testing.T) {
	dir := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(dir)
	if err := LoadConfig(); err == nil {
		t.Fatal("expected error for missing config")
	}
}

func TestParseFlag(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = []string{"cmd", "-loglevel=DEBUG"}
	ParseFlag()
	if Loglevel != "DEBUG" {
		t.Fatalf("expected DEBUG got %s", Loglevel)
	}
}
