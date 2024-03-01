package adapters

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func NewConfig() (*Config, error) {
	var config Config
	err := cleanenv.ReadConfig("config.yaml", &config)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
		return nil, err
	}
	return &config, err
}
