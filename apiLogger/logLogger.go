package apilogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)



func CreateLogger() *zap.Logger {
	logPath     := "../logs/log.json"
	//logPath     := "logs/log.json"

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
	return apiLogger

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