package utils

import (
	"log/slog"
)

func NewLogger() *slog.Logger {
	// options := &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}
	//handler := slog.NewTextHandler(os.Stdout, options)
	//handler := slog.NewJSONHandler(os.Stdout, options)

	//logger := slog.New(handler)
	logger := slog.Default()
	slog.SetDefault(logger)

	return logger
}
