package store

type Config struct {
	DataBaseUrl string `yaml:"DataBaseUrl"`
}

func NewConfig() *Config{
	return &Config{}
}