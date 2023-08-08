package envirements

import "os"

type EnvConfig struct {
	EnvRunAddr   string
	EnvResultURL string
}

// ParseEnv Get system environments
func ParseEnv() *EnvConfig {
	en := &EnvConfig{}
	en.EnvRunAddr = os.Getenv("RUN_ADDR")
	en.EnvResultURL = os.Getenv("SERVER_ADDRESS")
	return en
}
