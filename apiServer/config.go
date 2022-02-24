package apiServer

import (
	"GoTogether/apiLogger"
	"GoTogether/store"
)

type Config struct {

		BindAdrr string `yaml:"adrr"`
		Logger *apilogger.LogConfig
	    Store  *store.Config

}

func NewConfig() *Config  {
  return &Config{
  	BindAdrr: ":8080",
    Store: store.NewConfig(),
    Logger: apilogger.NewConfig(),

  }
}




