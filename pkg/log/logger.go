package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[0;36m"
	colorYellow = "\033[0;33m"
	colorRed    = "\033[0;31m"
)

var logger *slog.Logger

type CustomTextHandler struct {
	level slog.Level
}

func NewCustomTextHandler(level slog.Level) *CustomTextHandler {
	return &CustomTextHandler{level: level}
}

func (h *CustomTextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *CustomTextHandler) Handle(ctx context.Context, r slog.Record) error {
	timestamp := r.Time.Format("2006-01-02T15:04:05.000-07:00")
	level := r.Level.String()
	msg := r.Message

	color := colorForLevel(r.Level)

	logMessage := fmt.Sprintf("%s[%s] %s: %s%s", color, timestamp, level, msg, colorReset)

	fmt.Fprintln(os.Stdout, logMessage)

	return nil
}

func (h *CustomTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *CustomTextHandler) WithGroup(name string) slog.Handler {
	return h
}

func colorForLevel(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return colorCyan
	case slog.LevelWarn:
		return colorYellow
	case slog.LevelError:
		return colorRed
	default:
		return colorReset
	}
}

func Init(env string) {
	var logLevel slog.Level

	switch env {
	case "local":
		handler := NewCustomTextHandler(slog.LevelDebug)
		logger = slog.New(handler)
	default:
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})
		logger = slog.New(handler)
	}

	slog.SetDefault(logger)
}

func Debug(msg string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(msg, args...))
}

func Info(msg string, args ...interface{}) {
	logger.Info(fmt.Sprintf(msg, args...))
}

func Warn(msg string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(msg, args...))
}

func Error(msg string, args ...interface{}) {
	logger.Error(fmt.Sprintf(msg, args...))
}

func Fatal(msg string, args ...interface{}) {
	logger.Error(fmt.Sprintf(msg, args...))
	os.Exit(1)
}
