package config

import (
	flag "github.com/spf13/pflag"
)

// Config структура для хранения настроек
type Config struct {
	RunAddr   string
	ResultURL string
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags(cfg *Config) {
	//fs := flag.NewFlagSet("myflags", flag.ExitOnError)
	flag.StringVarP(&cfg.RunAddr, "a", "a", "localhost:8080", "Адрес запуска HTTP-сервера.")
	flag.StringVarP(&cfg.ResultURL, "b", "b", "http://localhost:8080/", "Адрес результирующего сокращённого URL.")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
