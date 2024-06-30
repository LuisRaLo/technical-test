package repositories

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLoggerConfig crea una nueva configuración de logger.
func NewLoggerConfig(level zap.AtomicLevel) zap.Config {
	cfg := zap.Config{
		Level:            level,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    "function",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	return cfg
}

// NewLogger crea un nuevo logger.
func NewLogger() (*zap.SugaredLogger, error) {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	ll := os.Getenv("LOG_LEVEL")
	if ll != "" {
		switch strings.ToLower(ll) {
		case "debug":
			level = zap.NewAtomicLevelAt(zap.DebugLevel)
		case "info":
			level = zap.NewAtomicLevelAt(zap.InfoLevel)
		case "warn":
			level = zap.NewAtomicLevelAt(zap.WarnLevel)
		case "error":
			level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		}
	}

	cfg := NewLoggerConfig(level)
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
