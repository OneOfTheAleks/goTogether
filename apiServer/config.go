package apiServer

import "GoTogether/store"

type Config struct {

		BindAdrr string `yaml:"adrr"`
	    Store  *store.Config

}

func NewConfig() *Config  {
  return &Config{
  	BindAdrr: ":8080",
    Store: store.NewConfig(),
  }
}