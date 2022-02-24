package apilogger

type LogConfig struct {
	logPath string `yaml:"logPath"`

}

func NewConfig() *LogConfig {
	return &LogConfig{
		logPath: "logs/log.json",
	}
}