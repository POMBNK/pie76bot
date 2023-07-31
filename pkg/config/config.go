package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"pie76bot/pkg/logger"
	"sync"
)

type Config struct {
	IsDebug bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	Storage struct {
		Type       string
		Postgresql Postgresql `json:"postgresql"`
	} `yaml:"storage"`
}

type Postgresql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var once sync.Once
var cfg = &Config{}

func GetCfg() *Config {
	once.Do(func() {
		logs := logger.GetLogger()
		logs.Info("Reading config...")
		err := cleanenv.ReadConfig("config.yml", cfg)
		if err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logs.Info(help)
			logs.Fatal(err)
		}
	})
	return cfg
}
