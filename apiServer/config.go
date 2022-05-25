package apiServer

import (
	"GoTogether/apiLogger"
)

type Config struct {

		BindAdrr string `yaml:"adrr"`
		Logger *apilogger.LogConfig
        DataBaseUrl string `yaml:"DataBaseUrl"`

}

func NewConfig() *Config  {
  return &Config{
  	BindAdrr: ":8080",
    Logger: apilogger.NewConfig(),

  }
}




