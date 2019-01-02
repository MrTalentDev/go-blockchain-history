package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Login            string `env:"LOGIN" envDefault:"bankex"`
	Password         string `env:"PASSWORD" envDefault:"default"`
	DatabaseIP       string `env:"IP" envDefault:"history.bankex.team"`
	DatabaseLogin    string `env:"LOGIN_DB" envDefault:"root"`
	DatabasePassword string `env:"PASSWORD_DB" envDefault:"default"`
	ContractAddress  string `env:"CONTRACT_ADDRESS " envDefault:"default"`
	ServerPort       string `env:"SERVER_PORT" envDefault:"8080"`
	KeyDB            string `env:"KEY_DB" envDefault:"test"`
}

var (
	ConfigInstance *Config
)

// GetConfig gets config instance.
func GetConfig() *Config {
	if ConfigInstance == nil {
		ConfigInstance = &Config{}
		err := env.Parse(ConfigInstance)
		if err != nil {
			log.Fatalf("error initializing config: %s", err)
		}
	}
	return ConfigInstance
}
