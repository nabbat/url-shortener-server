package config

import (
	"github.com/nabbat/url-shortener-server.git/internal/envirements"
	"github.com/nabbat/url-shortener-server.git/internal/flags"
	"strings"
)

// Config структура для хранения настроек
type Config struct {
	RunAddr   string
	ResultURL string
}

func SetEnv() *Config {
	c := &Config{}
	fl := flags.ParseFlags()
	en := envirements.ParseEnv()

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
	if !strings.HasPrefix(c.ResultURL, "http://") {
		c.ResultURL = "http://" + c.ResultURL
	}

	return c
}
