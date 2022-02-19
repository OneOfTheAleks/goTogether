package apiServer

type Config struct {

		BindAdrr string `yaml:"adrr"`

}

func NewConfig() *Config  {
  return &Config{
  	BindAdrr: ":8080",
  }
}