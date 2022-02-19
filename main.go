package main

import (
	"GoTogether/apiServer"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configPath string

func init(){
configPath = "configs/apiServer.yaml"
}

func main() {



	logger,_ := zap.NewProduction()
	defer logger.Sync()

	config := apiServer.NewConfig()
 //   C:=&config
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile,config)//C
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}





	s := apiServer.New(config)
	if err := s.Start(); err != nil {
		logger.Error(
			"Can't run server",
			zap.String("apiServer","Run"),
			zap.Error(err))
		//log.Fatal(err)
	}
}
