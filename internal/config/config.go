package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Server struct {
		Host               string        `yaml:"host"`
		Port               string        `yaml:"port"`
		ReadTimeout        time.Duration `yaml:"readTimeout"`
		WriteTimeout       time.Duration `yaml:"writeTimeout"`
		IdleTimeout        time.Duration `yaml:"idleTimeout"`
		MaxHeaderMegabytes int           `yaml:"maxHeaderBytes"`
	}

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbName"`
		SSLMode  string `yaml:"SSLMode"`
	}

	Jwt struct {
		Secret string `yaml:"secret"`
	}

	Config struct {
		Env      string   `yaml:"env" env-default:"local"`
		Origin   string   `yaml:"origin"`
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
		Jwt      Jwt      `yaml:"jwt"`
	}
)

func MustLoad() *Config {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local" // заменить на  local
	}

	configPath := fmt.Sprintf("./config/%s.yaml", appEnv)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("can not load config file: " + err.Error())
	}

	return &cfg
}
