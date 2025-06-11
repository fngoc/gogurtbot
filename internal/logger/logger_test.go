package logger

import "testing"

func TestInitialize(t *testing.T) {
	if err := Initialize("INFO"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if Log == nil {
		t.Fatal("log not initialized")
	}
}
