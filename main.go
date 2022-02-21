package main

import (
	"GoTogether/apiServer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var configPath string
var logPath    string

func init() {
	configPath = "configs/apiServer.yaml"
	logPath     = "logs/log.json"
	createLog(logPath)
}

func main() {

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
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)

	}
	//logger,_ := zap.NewProduction()
	//defer logger.Sync()

	config := apiServer.NewConfig()
	//   C:=&config

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, config) //C
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}


	s := apiServer.New(config)
	if err := s.Start(); err != nil {
	//	logger.Error(
	//		"Can't run server",
	//		zap.String("apiServer", "Run"),
	//		zap.Error(err))

			addLog(logger,
				"Can't run server",
				"apiServer",
				"Run",
				err)
	}
}

//---------------------------------

func addLog(log *zap.Logger, msg string, k string, v string, err error) {
	log.Error(
		msg,
		zap.String(k, v),
		zap.Error(err))
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
