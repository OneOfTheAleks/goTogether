package apilogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

type ApiLogger struct {
	logConfig *LogConfig
	Logger    *zap.Logger
}

func NewApiLogger(logConfig *LogConfig) *ApiLogger {
	return &ApiLogger{
		logConfig: logConfig,
	}
}

func (s *ApiLogger) Create() error {
	logPath :=s.logConfig.logPath
	createLog(logPath)
	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.ErrorLevel),
		OutputPaths: []string{logPath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	apiLogger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)

	}
	s.Logger = apiLogger
	return nil
}

func createLog(logP string) error {
	_, err := os.Open(logP)
	if err != nil {
		_, err := os.Create(logP)
		if err != nil {
			return err
		}
	}
	return nil
}