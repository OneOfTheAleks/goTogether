package main

import (
	"GoTogether/apiServer"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configPath string
//var logPath    string

func init() {
	configPath = "configs/apiServer.yaml"
//	logPath     = "logs/log.json"
//	createLog(logPath)
}

func main() {


	config := apiServer.NewConfig()
	//   C:=&config

	//c:= yaml.Marshal(config)
	d, err := yaml.Marshal(&config)

	fmt.Println(d)

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
	//	apiLogger.Error(
	//		"Can't run server",
	//		zap.String("apiServer", "Run"),
	//		zap.Error(err))

		//	addServerLog(config.Logger.,err)
		fmt.Println(err)
	}
}

//---------------------------------

func addLog(log *zap.Logger, msg string, k string, v string, err error) {
	log.Error(
		msg,
		zap.String(k, v),
		zap.Error(err))
}

func addServerLog(log *zap.Logger,err error)  {
	errString := "Api server error"
	log.Error(
		errString,
		zap.Error(err),
		)

}

