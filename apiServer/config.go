package apiServer

type Config struct {
	BindAdrr string
}

func NewConfig() *Config  {
  return &Config{
  	BindAdrr: ":8080",
  }
}