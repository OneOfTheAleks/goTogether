package main

import (
	"GoTogether/apiServer"
	"go.uber.org/zap"
)

var configPath string

func init(){

}

func main() {

	logger,_ := zap.NewProduction()
	defer logger.Sync()

	config := apiServer.NewConfig()
	s := apiServer.New(config)
	if err := s.Start(); err != nil {
		logger.Info("Can't run server")
		//log.Fatal(err)
	}
}
