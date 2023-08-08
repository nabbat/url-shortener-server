package envirements

import "os"

type EnvConfig struct {
	EnvRunAddr   string
	EnvResultURL string
}

// ParseEnv Get system environments
func ParseEnv() *EnvConfig {
	env := &EnvConfig{}
	env.EnvRunAddr = os.Getenv("RUN_ADDR")
	env.EnvResultURL = os.Getenv("SERVER_ADDRESS")
	return env
}
