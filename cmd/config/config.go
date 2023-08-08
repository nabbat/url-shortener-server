package config

import (
	"github.com/nabbat/url-shortener-server.git/internal/envirements"
	"github.com/nabbat/url-shortener-server.git/internal/flags"
)

// Config структура для хранения настроек
type Config struct {
	RunAddr   string
	ResultURL string
}

func SetEnv() *Config {
	fl := flags.ParseFlags()
	en := envirements.ParseEnv()
	c := &Config{}

	if en.EnvRunAddr != "" {
		c.RunAddr = en.EnvRunAddr
	} else {
		c.RunAddr = fl.RunAddr
	}

	if en.EnvResultURL != "" {
		c.ResultURL = en.EnvResultURL
	} else {
		c.ResultURL = fl.ResultURL
	}

	return c
}
