package main

import (
	"GoTogether/apiServer"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configPath string
var LogPath    string

func init() {
	configPath = "configs/apiServer.yaml"
	LogPath     = "logs/log.json"
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



	if err := apiServer.Start(config); err != nil {

		fmt.Println(err)
	}
	if err == nil  {
		fmt.Println("Server running")
	}
}

//---------------------------------



