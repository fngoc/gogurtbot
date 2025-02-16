package logger

import (
	"go.uber.org/zap"
)

var (
	Log      = zap.NewNop()
	loglevel = "INFO"
)

// Initialize инициализация логера
func Initialize() error {
	lvl, err := zap.ParseAtomicLevel(loglevel)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return err
	}
	Log = zl
	return nil
}
