package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP_server `yaml:"http_server"`
	RPS         `yaml:"rps"`
}

type HTTP_server struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	Idle_timeout time.Duration `yaml:"idle_timeout"`
}

type RPS struct {
	Requests int `yaml:"requests"`
	Seconds  int `yaml:"seconds"`
}

func MustLoad() *Config {
	var configPath string
	if len(os.Args) >= 2 && os.Args[1] != "" {
		configPath = os.Args[1]
	} else {
		configPath = "./config/config.yml"
		log.Println("set CONFIG_PATH=" + configPath)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
