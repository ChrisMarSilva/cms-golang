package logging

/*

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}


*/


package logger

import (
	"log"
	"os"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.Logger
}

func NewLogger() Logger {
	hostname, _ := os.Hostname()
	logConfiguration := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"host.name":           hostname,
			"service.instance.id": uuid.New().String(),
		},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			LevelKey:    "severity",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}

	logger, err := logConfiguration.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	return &zapLogger{logger: logger}
}

func (l *zapLogger) Info(msg string, fields ...Field) {
	f := l.toLoggerFields(fields...)
	l.logger.Info(msg, f...)
}

func (l *zapLogger) Error(msg string, fields ...Field) {
	f := l.toLoggerFields(fields...)
	l.logger.Error(msg, f...)
}

func (l *zapLogger) Debug(msg string, fields ...Field) {
	f := l.toLoggerFields(fields...)
	l.logger.Debug(msg, f...)
}

func (l *zapLogger) Warn(msg string, fields ...Field) {
	f := l.toLoggerFields(fields...)
	l.logger.Warn(msg, f...)
}

func (l *zapLogger) Fatal(msg string, fields ...Field) {
	f := l.toLoggerFields(fields...)
	l.logger.Fatal(msg, f...)
}

func (l *zapLogger) WithFields(fields ...Field) Logger {
	return &zapLogger{logger: l.logger.With(l.toLoggerFields(fields...)...)}
}

func (l *zapLogger) toLoggerFields(fields ...Field) []zapcore.Field {
	var loggerFields []zap.Field
	for _, field := range fields {
		loggerFields = append(loggerFields, zap.Any(field.Key, field.Value))
	}
	return loggerFields
}

package logger

type Level uint

const (
	InfoLevel Level = iota
	ErrorLevel
	WarnLevel
	DebugLevel
)

type Field struct {
	Key   string
	Value any
}

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	WithFields(fields ...Field) Logger
}