package config

import (
	"final-project-enigma-clean/util/helper"
	"fmt"
	"github.com/gookit/slog"
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	DbDriver string
}

type LoggerPath struct {
	FilePath string
}

type Config struct {
	*DbConfig
	*LoggerPath
}

func (c *Config) ReadConfig() error {
	//load config from env
	if err := helper.LoadEnv(); err != nil {
		return err
	}
	c.DbConfig = &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		DbDriver: os.Getenv("DB_DRIVER"),
	}

	//file config
	c.LoggerPath = &LoggerPath{
		FilePath: os.Getenv("LOGGER_FILE"),
	}

	//if .env is missing
	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.DbName == "" ||
		c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.DbDriver == "" {
		return fmt.Errorf("missing required environment variable")
	}

	slog.Infof("Connected to database %v", c.DbConfig.Host)
	fmt.Println("-------------------------------------")
	return nil

}

func NewDbConfig() (*Config, error) {
	cfg := &Config{}

	//if read config is failed
	if err := cfg.ReadConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read config %v", err.Error())
	}

	return cfg, nil
}
